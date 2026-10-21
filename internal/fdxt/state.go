package fdxt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"fdxt-poc/internal/crypto"
)

type clientStateFile struct {
	KW     string              `json:"kw"`
	KT     string              `json:"kt"`
	KX     string              `json:"kx"`
	KY     string              `json:"ky"`
	KZ     string              `json:"kz"`
	Count  map[string]Counter  `json:"count"`
	C1Pool []clientC1PoolEntry `json:"c1_pool,omitempty"`
}

type clientC1PoolKeyword struct {
	J uint64 `json:"j"`
	L string `json:"l"`
	T string `json:"t"`
}

type clientC1PoolEntry struct {
	ID       string                         `json:"id"`
	Keywords map[string]clientC1PoolKeyword `json:"keywords"`
}

type serverTEntry struct {
	Val   string `json:"val"`
	Alpha string `json:"alpha"`
}
type serverStateFile struct {
	// CDBX/CDBT keys and XSET values are arbitrary binary strings. They must
	// be base64-encoded before entering JSON; encoding raw binary Go strings
	// as JSON strings silently replaces invalid UTF-8 and breaks remote search.
	CDBX map[string]string       `json:"cdbx"`
	CDBT map[string]serverTEntry `json:"cdbt"`
	XSET []string                `json:"xset"`
}

// SaveClientState persists only client secrets/local counters. It is not
// sent to the server.
func (c *Client) SaveState(path string) error {
	s := clientStateFile{
		KW: b64(c.KW), KT: b64(c.KT), KX: b64(c.KX), KY: b64(c.KY), KZ: b64(c.KZ),
		Count:  make(map[string]Counter, len(c.Count)),
		C1Pool: make([]clientC1PoolEntry, 0, len(c.C1Pool)),
	}
	for k, v := range c.Count {
		s.Count[k] = *v
	}
	for _, ce := range c.C1Pool {
		pe := clientC1PoolEntry{ID: ce.ID, Keywords: make(map[string]clientC1PoolKeyword, len(ce.Keywords))}
		for kw, pk := range ce.Keywords {
			pe.Keywords[kw] = clientC1PoolKeyword{J: pk.J, L: b64(pk.L), T: b64(pk.T)}
		}
		s.C1Pool = append(s.C1Pool, pe)
	}
	return writeJSON(path, s)
}

func LoadClientState(path string, p *crypto.Ctx) (*Client, error) {
	var s clientStateFile
	if err := readJSON(path, &s); err != nil {
		return nil, err
	}
	c := &Client{P: p, Count: make(map[string]*Counter)}
	var err error
	if c.KW, err = unb64(s.KW); err != nil {
		return nil, fmt.Errorf("KW: %w", err)
	}
	if c.KT, err = unb64(s.KT); err != nil {
		return nil, fmt.Errorf("KT: %w", err)
	}
	if c.KX, err = unb64(s.KX); err != nil {
		return nil, fmt.Errorf("KX: %w", err)
	}
	if c.KY, err = unb64(s.KY); err != nil {
		return nil, fmt.Errorf("KY: %w", err)
	}
	if c.KZ, err = unb64(s.KZ); err != nil {
		return nil, fmt.Errorf("KZ: %w", err)
	}
	for k, v := range s.Count {
		vv := v
		c.Count[k] = &vv
	}
	for _, pe := range s.C1Pool {
		ce := C1PoolEntry{ID: pe.ID, Keywords: make(map[string]C1PoolKeyword, len(pe.Keywords))}
		for kw, pk := range pe.Keywords {
			l, err := unb64(pk.L)
			if err != nil {
				return nil, fmt.Errorf("c1 pool L: %w", err)
			}
			t, err := unb64(pk.T)
			if err != nil {
				return nil, fmt.Errorf("c1 pool T: %w", err)
			}
			ce.Keywords[kw] = C1PoolKeyword{J: pk.J, L: l, T: t}
		}
		c.C1Pool = append(c.C1Pool, ce)
	}
	return c, nil
}

// SaveServerState persists the server's encrypted database and XSET.
// PBC Zr alpha values are serialized as their canonical Bytes().
func (s *Server) SaveState(path string) error {
	out := serverStateFile{CDBX: make(map[string]string, len(s.CDBX)), CDBT: make(map[string]serverTEntry, len(s.CDBT))}
	for k, v := range s.CDBX {
		out.CDBX[b64([]byte(k))] = b64(v)
	}
	for k, v := range s.CDBT {
		out.CDBT[b64([]byte(k))] = serverTEntry{Val: b64(v.Val), Alpha: b64(v.Alpha.Bytes())}
	}
	for k := range s.XSET {
		out.XSET = append(out.XSET, b64([]byte(k)))
	}
	sort.Strings(out.XSET)
	return writeJSON(path, out)
}

func LoadServerState(path string, p *crypto.Ctx) (*Server, error) {
	var in serverStateFile
	if err := readJSON(path, &in); err != nil {
		return nil, err
	}
	s := NewServer(p)
	for k, v := range in.CDBX {
		key, err := unb64(k)
		if err != nil {
			return nil, fmt.Errorf("CDBX key: %w", err)
		}
		b, err := unb64(v)
		if err != nil {
			return nil, err
		}
		s.CDBX[string(key)] = b
	}
	for k, v := range in.CDBT {
		key, err := unb64(k)
		if err != nil {
			return nil, fmt.Errorf("CDBT key: %w", err)
		}
		val, err := unb64(v.Val)
		if err != nil {
			return nil, err
		}
		ab, err := unb64(v.Alpha)
		if err != nil {
			return nil, err
		}
		alpha := p.Pairing.NewZr()
		alpha.SetBytes(ab)
		s.CDBT[string(key)] = CDBTEntry{Val: val, Alpha: alpha}
	}
	for _, k := range in.XSET {
		key, err := unb64(k)
		if err != nil {
			return nil, fmt.Errorf("XSET key: %w", err)
		}
		s.XSET[string(key)] = true
	}
	return s, nil
}

func b64(b []byte) string            { return base64.StdEncoding.EncodeToString(b) }
func unb64(s string) ([]byte, error) { return base64.StdEncoding.DecodeString(s) }

func writeJSON(path string, v any) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepathDir(path), 0700); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}
func readJSON(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
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
