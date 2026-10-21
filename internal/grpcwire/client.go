package grpcwire

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"fmt"
	"net/http"
	"os"
	"time"

	"fdxt-poc/internal/fdxt"
	"fdxt-poc/internal/grpcproto"
)

type Client struct {
	URL  string
	HTTP *http.Client
}

func NewClient(url, certPath string) (*Client, error) {
	pemData, err := os.ReadFile(certPath)
	if err != nil {
		return nil, err
	}
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(pemData) {
		return nil, fmt.Errorf("failed to load server certificate")
	}
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{RootCAs: pool, MinVersion: tls.VersionTLS12, NextProtos: []string{"h2"}},
		ForceAttemptHTTP2: true,
	}
	return &Client{URL: url, HTTP: &http.Client{Transport: tr, Timeout: 60 * time.Second}}, nil
}

func (c *Client) Search(parts fdxt.SearchRequestParts) ([]fdxt.SearchOutput, error) {
	req := &grpcproto.SearchRequest{}
	for _, x := range parts.TKL {
		req.TKL = append(req.TKL, &grpcproto.TKLEntry{L: x.L, T: x.T})
	}
	for _, row := range parts.Rows {
		r := &grpcproto.XTKRow{Addr: row.Addr}
		for kw, e := range row.XTK {
			r.XTK = append(r.XTK, &grpcproto.XTKEntry{Keyword: kw, Element: e.Bytes()})
		}
		req.Rows = append(req.Rows, r)
	}
	req.RestKeywords = append([]string(nil), parts.RestKeywords...)

	body := req.Marshal()
	framed := make([]byte, 5+len(body))
	copy(framed[5:], body)
	binary.BigEndian.PutUint32(framed[1:], uint32(len(body)))
	httpReq, err := http.NewRequest(http.MethodPost, c.URL, bytes.NewReader(framed))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/grpc")
	httpReq.Header.Set("TE", "trailers")
	resp, err := c.HTTP.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gRPC HTTP status %s", resp.Status)
	}
	msg, err := readFrame(resp.Body)
	if err != nil {
		return nil, err
	}
	if st := resp.Trailer.Get("grpc-status"); st != "" && st != "0" {
		return nil, fmt.Errorf("gRPC status %s: %s", st, resp.Trailer.Get("grpc-message"))
	}
	var out grpcproto.SearchResponse
	if err := out.Unmarshal(msg); err != nil {
		return nil, err
	}
	res := make([]fdxt.SearchOutput, 0, len(out.Outputs))
	for _, x := range out.Outputs {
		res = append(res, fdxt.SearchOutput{Val: x.Val, Cnt: int(x.Cnt), Addr: append([]byte(nil), x.Addr...)})
	}
	return res, nil
}

// SearchWireBytes returns the exact protobuf message sizes (including the
// 5-byte gRPC frame prefix) that the current transport would send for a
// request/response pair. It is used by the experiment runner so communication
// cost is derived from the same wire representation as the real gRPC path.
func SearchWireBytes(parts fdxt.SearchRequestParts, res []fdxt.SearchOutput) (requestBytes, responseBytes int) {
	req := &grpcproto.SearchRequest{}
	for _, x := range parts.TKL {
		req.TKL = append(req.TKL, &grpcproto.TKLEntry{L: x.L, T: x.T})
	}
	for _, row := range parts.Rows {
		r := &grpcproto.XTKRow{Addr: row.Addr}
		for kw, e := range row.XTK {
			r.XTK = append(r.XTK, &grpcproto.XTKEntry{Keyword: kw, Element: e.Bytes()})
		}
		req.Rows = append(req.Rows, r)
	}
	req.RestKeywords = append([]string(nil), parts.RestKeywords...)

	out := &grpcproto.SearchResponse{}
	for _, x := range res {
		out.Outputs = append(out.Outputs, &grpcproto.SearchOutput{Val: x.Val, Cnt: uint32(x.Cnt), Addr: x.Addr})
	}
	return 5 + len(req.Marshal()), 5 + len(out.Marshal())
}
