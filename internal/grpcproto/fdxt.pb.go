// Code generated in the style of protoc output for this POC's fixed
// wire schema. The repository does not require protoc at runtime; these
// types implement the protobuf wire encoding used by the gRPC transport.
package grpcproto

import (
	"encoding/binary"
	"fmt"
)

type TKLEntry struct {
	L []byte
	T []byte
}
type XTKEntry struct {
	Keyword string
	Element []byte
}
type XTKRow struct {
	Addr []byte
	XTK  []*XTKEntry
}
type SearchRequest struct {
	TKL          []*TKLEntry
	Rows         []*XTKRow
	RestKeywords []string
	C1State      []byte
	C1MaxCover   uint32
}
type SearchOutput struct {
	Val  []byte
	Cnt  uint32
	Addr []byte
}
type SearchResponse struct{ Outputs []*SearchOutput }

func appendVarint(dst []byte, field int, v uint64) []byte {
	dst = append(dst, byte(field<<3|0))
	var b [10]byte
	n := binary.PutUvarint(b[:], v)
	return append(dst, b[:n]...)
}
func appendBytes(dst []byte, field int, v []byte) []byte {
	if len(v) == 0 {
		return dst
	}
	dst = append(dst, byte(field<<3|2))
	var b [10]byte
	n := binary.PutUvarint(b[:], uint64(len(v)))
	dst = append(dst, b[:n]...)
	return append(dst, v...)
}
func appendString(dst []byte, field int, v string) []byte { return appendBytes(dst, field, []byte(v)) }
func appendMsg(dst []byte, field int, msg []byte) []byte  { return appendBytes(dst, field, msg) }

func (m *TKLEntry) Marshal() []byte {
	var b []byte
	b = appendBytes(b, 1, m.L)
	b = appendBytes(b, 2, m.T)
	return b
}
func (m *XTKEntry) Marshal() []byte {
	var b []byte
	b = appendString(b, 1, m.Keyword)
	b = appendBytes(b, 2, m.Element)
	return b
}
func (m *XTKRow) Marshal() []byte {
	var b []byte
	b = appendBytes(b, 1, m.Addr)
	for _, x := range m.XTK {
		if x != nil {
			b = appendMsg(b, 2, x.Marshal())
		}
	}
	return b
}
func (m *SearchRequest) Marshal() []byte {
	var b []byte
	for _, x := range m.TKL {
		if x != nil {
			b = appendMsg(b, 1, x.Marshal())
		}
	}
	for _, x := range m.Rows {
		if x != nil {
			b = appendMsg(b, 2, x.Marshal())
		}
	}
	for _, x := range m.RestKeywords {
		b = appendString(b, 3, x)
	}
	if len(m.C1State) != 0 {
		b = appendBytes(b, 4, m.C1State)
	}
	if m.C1MaxCover != 0 {
		b = appendVarint(b, 5, uint64(m.C1MaxCover))
	}
	return b
}
func (m *SearchOutput) Marshal() []byte {
	var b []byte
	b = appendBytes(b, 1, m.Val)
	if m.Cnt != 0 {
		b = appendVarint(b, 2, uint64(m.Cnt))
	}
	b = appendBytes(b, 3, m.Addr)
	return b
}
func (m *SearchResponse) Marshal() []byte {
	var b []byte
	for _, x := range m.Outputs {
		if x != nil {
			b = appendMsg(b, 1, x.Marshal())
		}
	}
	return b
}

func readField(data []byte, off *int) (int, int, []byte, error) {
	if *off >= len(data) {
		return 0, 0, nil, fmt.Errorf("eof")
	}
	tag, n := binary.Uvarint(data[*off:])
	if n <= 0 {
		return 0, 0, nil, fmt.Errorf("bad tag")
	}
	*off += n
	field := int(tag >> 3)
	wt := int(tag & 7)
	switch wt {
	case 0:
		v, n := binary.Uvarint(data[*off:])
		if n <= 0 {
			return 0, 0, nil, fmt.Errorf("bad varint")
		}
		*off += n
		var tmp [10]byte
		nn := binary.PutUvarint(tmp[:], v)
		return field, wt, tmp[:nn], nil
	case 2:
		l, n := binary.Uvarint(data[*off:])
		if n <= 0 {
			return 0, 0, nil, fmt.Errorf("bad length")
		}
		*off += n
		if l > uint64(len(data)-*off) {
			return 0, 0, nil, fmt.Errorf("length exceeds message")
		}
		v := data[*off : *off+int(l)]
		*off += int(l)
		return field, wt, v, nil
	default:
		return 0, 0, nil, fmt.Errorf("unsupported wire type %d", wt)
	}
}
func varintBytes(v []byte) uint64 { x, _ := binary.Uvarint(v); return x }

func (m *TKLEntry) Unmarshal(data []byte) error {
	for off := 0; off < len(data); {
		f, wt, v, e := readField(data, &off)
		if e != nil {
			return e
		}
		if wt != 2 {
			return fmt.Errorf("TKLEntry field %d has wrong wire type", f)
		}
		switch f {
		case 1:
			m.L = append([]byte(nil), v...)
		case 2:
			m.T = append([]byte(nil), v...)
		}
	}
	return nil
}
func (m *XTKEntry) Unmarshal(data []byte) error {
	for off := 0; off < len(data); {
		f, wt, v, e := readField(data, &off)
		if e != nil {
			return e
		}
		if wt != 2 {
			return fmt.Errorf("XTKEntry field %d wrong wire type", f)
		}
		switch f {
		case 1:
			m.Keyword = string(v)
		case 2:
			m.Element = append([]byte(nil), v...)
		}
	}
	return nil
}
func (m *XTKRow) Unmarshal(data []byte) error {
	for off := 0; off < len(data); {
		f, wt, v, e := readField(data, &off)
		if e != nil {
			return e
		}
		if wt != 2 {
			return fmt.Errorf("XTKRow field %d wrong wire type", f)
		}
		switch f {
		case 1:
			m.Addr = append([]byte(nil), v...)
		case 2:
			x := new(XTKEntry)
			if e := x.Unmarshal(v); e != nil {
				return e
			}
			m.XTK = append(m.XTK, x)
		}
	}
	return nil
}
func (m *SearchRequest) Unmarshal(data []byte) error {
	for off := 0; off < len(data); {
		f, wt, v, e := readField(data, &off)
		if e != nil {
			return e
		}
		switch f {
		case 1:
			if wt != 2 {
				return fmt.Errorf("tkl wire")
			}
			x := new(TKLEntry)
			if e := x.Unmarshal(v); e != nil {
				return e
			}
			m.TKL = append(m.TKL, x)
		case 2:
			if wt != 2 {
				return fmt.Errorf("row wire")
			}
			x := new(XTKRow)
			if e := x.Unmarshal(v); e != nil {
				return e
			}
			m.Rows = append(m.Rows, x)
		case 3:
			if wt != 2 {
				return fmt.Errorf("keyword wire")
			}
			m.RestKeywords = append(m.RestKeywords, string(v))
		case 4:
			if wt != 2 {
				return fmt.Errorf("c1 state wire")
			}
			m.C1State = append([]byte(nil), v...)
		case 5:
			if wt != 0 {
				return fmt.Errorf("c1 max cover wire")
			}
			m.C1MaxCover = uint32(varintBytes(v))
		}
	}
	return nil
}
func (m *SearchOutput) Unmarshal(data []byte) error {
	for off := 0; off < len(data); {
		f, wt, v, e := readField(data, &off)
		if e != nil {
			return e
		}
		switch f {
		case 1:
			if wt != 2 {
				return fmt.Errorf("val wire")
			}
			m.Val = append([]byte(nil), v...)
		case 2:
			if wt != 0 {
				return fmt.Errorf("cnt wire")
			}
			m.Cnt = uint32(varintBytes(v))
		case 3:
			if wt != 2 {
				return fmt.Errorf("addr wire")
			}
			m.Addr = append([]byte(nil), v...)
		}
	}
	return nil
}
func (m *SearchResponse) Unmarshal(data []byte) error {
	for off := 0; off < len(data); {
		f, wt, v, e := readField(data, &off)
		if e != nil {
			return e
		}
		if f == 1 {
			if wt != 2 {
				return fmt.Errorf("output wire")
			}
			x := new(SearchOutput)
			if e := x.Unmarshal(v); e != nil {
				return e
			}
			m.Outputs = append(m.Outputs, x)
		}
	}
	return nil
}
