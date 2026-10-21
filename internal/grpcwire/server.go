package grpcwire

import (
	"encoding/binary"
	"io"
	"net"

	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"sync"
	"time"

	fdxtcrypto "fdxt-poc/internal/crypto"
	"fdxt-poc/internal/fdxt"
	"fdxt-poc/internal/grpcproto"
	"github.com/Nik-U/pbc"
)

// Server implements the FDXT gRPC wire service using the standard-library
// HTTP/2 stack. The endpoint is wire-compatible with gRPC: HTTP/2,
// application/grpc content type, 5-byte protobuf message framing, and
// grpc-status trailers.
type Server struct {
	FDXT      *fdxt.Server
	StatePath string
	mu        sync.Mutex
}

func (s *Server) Serve(addr, certPath, keyPath string) error {
	if err := ensureCertificate(certPath, keyPath); err != nil {
		return err
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/fdxt.FDXT/Search", s.searchHandler)
	h := &http.Server{Addr: addr, Handler: mux, ReadHeaderTimeout: 10 * time.Second}
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return err
	}
	h.TLSConfig = &tls.Config{Certificates: []tls.Certificate{cert}, MinVersion: tls.VersionTLS12, NextProtos: []string{"h2"}}
	fmt.Printf("[SERVER] FDXT gRPC endpoint listening on https://%s/fdxt.FDXT/Search\n", addr)
	return h.ListenAndServeTLS("", "")
}

func (s *Server) searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" || r.Header.Get("Content-Type") != "application/grpc" {
		w.Header().Set("grpc-status", "12")
		http.Error(w, "unimplemented", http.StatusNotImplemented)
		return
	}
	body, err := readFrame(r.Body)
	if err != nil {
		writeGRPCError(w, 3, err)
		return
	}
	var req grpcproto.SearchRequest
	if err := req.Unmarshal(body); err != nil {
		writeGRPCError(w, 3, err)
		return
	}
	parts, err := fromProto(req, s.FDXT.P)
	if err != nil {
		writeGRPCError(w, 3, err)
		return
	}
	s.mu.Lock()
	res := s.FDXT.ProcessQuery(fdxt.SearchRequestParts{TKL: parts.tkl, Rows: parts.rows, RestKeywords: parts.rest})
	if s.StatePath != "" {
		err = s.FDXT.SaveState(s.StatePath)
	}
	s.mu.Unlock()
	if err != nil {
		writeGRPCError(w, 13, err)
		return
	}
	out := &grpcproto.SearchResponse{}
	for _, x := range res {
		out.Outputs = append(out.Outputs, &grpcproto.SearchOutput{Val: x.Val, Cnt: uint32(x.Cnt), Addr: x.Addr})
	}
	writeGRPCResponse(w, out.Marshal())
}

type wireParts struct {
	tkl  []fdxt.TKLEntry
	rows []fdxt.XTKRow
	rest []string
}

func fromProto(req grpcproto.SearchRequest, p *fdxtcrypto.Ctx) (wireParts, error) {
	var out wireParts
	for _, x := range req.TKL {
		out.tkl = append(out.tkl, fdxt.TKLEntry{L: x.L, T: x.T})
	}
	for _, row := range req.Rows {
		r := fdxt.XTKRow{Addr: row.Addr, XTK: make(map[string]*pbc.Element)}
		for _, x := range row.XTK {
			e := p.Pairing.NewG1()
			if len(x.Element) == 0 {
				return out, fmt.Errorf("empty XTK element for %q", x.Keyword)
			}
			e.SetBytes(x.Element)
			r.XTK[x.Keyword] = e
		}
		out.rows = append(out.rows, r)
	}
	out.rest = append([]string(nil), req.RestKeywords...)
	return out, nil
}

func readFrame(r io.Reader) ([]byte, error) {
	var hdr [5]byte
	if _, err := io.ReadFull(r, hdr[:]); err != nil {
		return nil, err
	}
	if hdr[0] != 0 {
		return nil, fmt.Errorf("compressed gRPC messages are unsupported")
	}
	n := int(binary.BigEndian.Uint32(hdr[1:]))
	if n < 0 || n > 64<<20 {
		return nil, fmt.Errorf("invalid message length %d", n)
	}
	b := make([]byte, n)
	_, err := io.ReadFull(r, b)
	return b, err
}
func writeGRPCResponse(w http.ResponseWriter, msg []byte) {
	w.Header().Set("Content-Type", "application/grpc")
	w.Header().Add("Trailer", "grpc-status")
	w.WriteHeader(http.StatusOK)
	var hdr [5]byte
	binary.BigEndian.PutUint32(hdr[1:], uint32(len(msg)))
	_, _ = w.Write(hdr[:])
	_, _ = w.Write(msg)
	w.Header().Set("grpc-status", "0")
}
func writeGRPCError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/grpc")
	w.Header().Add("Trailer", "grpc-status")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("grpc-status", fmt.Sprintf("%d", status))
	var hdr [5]byte
	_, _ = w.Write(hdr[:])
	_ = err
}

func ensureCertificate(certPath, keyPath string) error {
	if _, err := os.Stat(certPath); err == nil {
		if _, e := os.Stat(keyPath); e == nil {
			return nil
		}
	}
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	serial, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 120))
	if err != nil {
		return err
	}
	tmpl := x509.Certificate{SerialNumber: serial, Subject: pkix.Name{CommonName: "fdxt-localhost"}, NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(365 * 24 * time.Hour), KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, DNSNames: []string{"localhost"}, IPAddresses: []net.IP{net.ParseIP("127.0.0.1")}}
	der, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, &key.PublicKey, key)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepathDir(certPath), 0700); err != nil {
		return err
	}
	cf, err := os.Create(certPath)
	if err != nil {
		return err
	}
	defer cf.Close()
	if err := pem.Encode(cf, &pem.Block{Type: "CERTIFICATE", Bytes: der}); err != nil {
		return err
	}
	kf, err := os.Create(keyPath)
	if err != nil {
		return err
	}
	defer kf.Close()
	return pem.Encode(kf, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
}
func filepathDir(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' || path[i] == '\\' {
			if i == 0 {
				return "."
			}
			return path[:i]
		}
	}
	return "."
}
