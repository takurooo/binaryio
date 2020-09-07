package binaryio

import (
	"fmt"
	"io"
)

// Writer ...
type Writer struct {
	io.WriterAt
	offset int64
	err    error
	b8     []byte
	b16    []byte
	b24    []byte
	b32    []byte
	b64    []byte
}

// NewWriter ...
func NewWriter(w io.WriterAt) (br *Writer) {
	br = &Writer{
		w,               // io.WriterAt
		0,               // offset
		nil,             // err
		make([]byte, 1), // b8
		make([]byte, 2), // b16
		make([]byte, 3), // b24
		make([]byte, 4), // b32
		make([]byte, 8), // b64
	}
	return br
}

func (bw *Writer) writeBytes(p []byte) (n int) {
	n, bw.err = bw.WriteAt(p, bw.offset)
	bw.offset += int64(n)
	return n
}

func (bw *Writer) setErr(err error) {
	bw.err = err
}

// Err ...
func (bw *Writer) Err() error {
	return bw.err
}

// GetOffset ...
func (bw *Writer) GetOffset() int64 {
	return bw.offset
}

// SetOffset ...
func (bw *Writer) SetOffset(offset int64) {
	bw.offset = offset
}

// WriteRaw ...
func (bw *Writer) WriteRaw(p []byte) int {
	if bw.err != nil {
		return 0
	}
	return bw.writeBytes(p)
}

// WriteI8 ...
func (bw *Writer) WriteI8(v int8) int {
	if bw.err != nil {
		return 0
	}
	bw.b8[0] = byte(v)
	return bw.writeBytes(bw.b8)
}

// WriteI16 ...
func (bw *Writer) WriteI16(v int16, e Endian) int {
	if bw.err != nil {
		return 0
	}
	return bw.WriteU16(uint16(v), e)
}

// WriteI24 ...
func (bw *Writer) WriteI24(v int32, e Endian) int {
	if bw.err != nil {
		return 0
	}
	return bw.WriteU24(uint32(v), e)
}

// WriteI32 ...
func (bw *Writer) WriteI32(v int32, e Endian) int {
	if bw.err != nil {
		return 0
	}
	return bw.WriteU32(uint32(v), e)
}

// WriteI64 ...
func (bw *Writer) WriteI64(v int64, e Endian) int {
	if bw.err != nil {
		return 0
	}
	return bw.WriteU64(uint64(v), e)
}

// WriteU8 ...
func (bw *Writer) WriteU8(v uint8) int {
	if bw.err != nil {
		return 0
	}
	bw.b8[0] = v
	return bw.writeBytes(bw.b8)
}

// WriteU16 ...
func (bw *Writer) WriteU16(v uint16, e Endian) int {
	if bw.err != nil {
		return 0
	}

	if e == LittleEndian {
		bw.b16[0] = byte(v)
		bw.b16[1] = byte(v >> 8)
	} else {
		bw.b16[0] = byte(v >> 8)
		bw.b16[1] = byte(v)
	}

	return bw.writeBytes(bw.b16)
}

// WriteU24 ...
func (bw *Writer) WriteU24(v uint32, e Endian) int {
	if bw.err != nil {
		return 0
	}

	if e == LittleEndian {
		bw.b24[0] = byte(v)
		bw.b24[1] = byte(v >> 8)
		bw.b24[2] = byte(v >> 16)
	} else {
		bw.b24[0] = byte(v >> 16)
		bw.b24[1] = byte(v >> 8)
		bw.b24[2] = byte(v)
	}

	return bw.writeBytes(bw.b24)
}

// WriteU32 ...
func (bw *Writer) WriteU32(v uint32, e Endian) int {
	if bw.err != nil {
		return 0
	}

	if e == LittleEndian {
		bw.b32[0] = byte(v)
		bw.b32[1] = byte((v >> 8))
		bw.b32[2] = byte((v >> 16))
		bw.b32[3] = byte((v >> 24))
	} else {
		bw.b32[0] = byte((v >> 24))
		bw.b32[1] = byte((v >> 16))
		bw.b32[2] = byte((v >> 8))
		bw.b32[3] = byte(v)
	}

	return bw.writeBytes(bw.b32)
}

// WriteU64 ...
func (bw *Writer) WriteU64(v uint64, e Endian) int {
	if bw.err != nil {
		return 0
	}

	if e == LittleEndian {
		bw.b64[0] = byte(v)
		bw.b64[1] = byte(v >> 8)
		bw.b64[2] = byte(v >> 16)
		bw.b64[3] = byte(v >> 24)
		bw.b64[4] = byte(v >> 32)
		bw.b64[5] = byte(v >> 40)
		bw.b64[6] = byte(v >> 48)
		bw.b64[7] = byte(v >> 56)
	} else {
		bw.b64[0] = byte(v >> 56)
		bw.b64[1] = byte(v >> 48)
		bw.b64[2] = byte(v >> 40)
		bw.b64[3] = byte(v >> 32)
		bw.b64[4] = byte(v >> 24)
		bw.b64[5] = byte(v >> 16)
		bw.b64[6] = byte(v >> 8)
		bw.b64[7] = byte(v)
	}

	return bw.writeBytes(bw.b64)
}

// WriteS8 ...
func (bw *Writer) WriteS8(s string) int {
	if bw.err != nil {
		return 0
	}
	return bw.WriteU8(uint8(s[0]))
}

// WriteS16 ...
func (bw *Writer) WriteS16(s string, e Endian) int {
	if bw.err != nil {
		return 0
	}
	var v uint16
	v = uint16(s[0])<<8 |
		uint16(s[1])

	return bw.WriteU16(v, e)
}

// WriteS24 ...
func (bw *Writer) WriteS24(s string, e Endian) int {
	if bw.err != nil {
		return 0
	}
	var v uint32
	v = uint32(s[0])<<16 |
		uint32(s[1])<<8 |
		uint32(s[2])

	return bw.WriteU24(v, e)
}

// WriteS32 ...
func (bw *Writer) WriteS32(s string, e Endian) int {
	if bw.err != nil {
		return 0
	}
	var v uint32
	v = uint32(s[0])<<24 |
		uint32(s[1])<<16 |
		uint32(s[2])<<8 |
		uint32(s[3])

	return bw.WriteU32(v, e)
}

// WriteS64 ...
func (bw *Writer) WriteS64(s string, e Endian) int {
	if bw.err != nil {
		return 0
	}
	var v uint64
	v = uint64(s[0])<<56 |
		uint64(s[1])<<48 |
		uint64(s[2])<<40 |
		uint64(s[3])<<32 |
		uint64(s[4])<<24 |
		uint64(s[5])<<16 |
		uint64(s[6])<<8 |
		uint64(s[7])
	return bw.WriteU64(v, e)
}

// WriteX ...
func (bw *Writer) WriteX(e Endian, chainData ...interface{}) int {

	var n int
	for _, d := range chainData {
		switch v := d.(type) {
		case []int8:
			for _, x := range v {
				n += bw.WriteI8(x)
			}
		case []int16:
			for _, x := range v {
				n += bw.WriteI16(x, e)
			}
		case []int32:
			for _, x := range v {
				n += bw.WriteI32(x, e)
			}
		case []int64:
			for _, x := range v {
				n += bw.WriteI64(x, e)
			}
		case []uint8:
			for _, x := range v {
				n += bw.WriteU8(x)
			}
		case []uint16:
			for _, x := range v {
				n += bw.WriteU16(x, e)
			}
		case []uint32:
			for _, x := range v {
				n += bw.WriteU32(x, e)
			}
		case []uint64:
			for _, x := range v {
				n += bw.WriteU64(x, e)
			}
		case int8:
			n += bw.WriteI8(v)
		case int16:
			n += bw.WriteI16(v, e)
		case int32:
			n += bw.WriteI32(v, e)
		case int64:
			n += bw.WriteI64(v, e)
		case uint8:
			n += bw.WriteU8(v)
		case uint16:
			n += bw.WriteU16(v, e)
		case uint32:
			n += bw.WriteU32(v, e)
		case uint64:
			n += bw.WriteU64(v, e)
		case *int8:
			n += bw.WriteI8(*v)
		case *int16:
			n += bw.WriteI16(*v, e)
		case *int32:
			n += bw.WriteI32(*v, e)
		case *int64:
			n += bw.WriteI64(*v, e)
		case *uint8:
			n += bw.WriteU8(*v)
		case *uint16:
			n += bw.WriteU16(*v, e)
		case *uint32:
			n += bw.WriteU32(*v, e)
		case *uint64:
			n += bw.WriteU64(*v, e)
		default:
			panic(fmt.Errorf("not supported type %T", v))
		}
	}

	return n

}
