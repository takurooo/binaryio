package binaryio

import (
	"io"
)

// Writer ...
type Writer struct {
	io.WriterAt
	offset int64
	err    error
}

// NewWriter ...
func NewWriter(w io.WriterAt) (br *Writer) {
	br = &Writer{w, 0, nil}
	return br
}

func (bw *Writer) writeBytes(p []byte) (n int) {
	offset := bw.offset

	n, err := bw.WriteAt(p, offset)
	bw.setErr(err)

	defer func() {
		bw.offset += int64(n)
	}()

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
func (bw *Writer) WriteRaw(p []byte) (n int) {
	if bw.err != nil {
		return 0
	}
	n = bw.writeBytes(p)
	return n
}

// WriteI8 ...
func (bw *Writer) WriteI8(v int8) (n int) {
	if bw.err != nil {
		return 0
	}
	b := make([]byte, 1)
	b[0] = byte(v)
	n = bw.writeBytes(b)

	return n
}

// WriteI16 ...
func (bw *Writer) WriteI16(v int16, e Endian) (n int) {
	if bw.err != nil {
		return 0
	}
	n = bw.WriteU16(uint16(v), e)
	return n
}

// WriteI24 ...
func (bw *Writer) WriteI24(v int32, e Endian) (n int) {
	if bw.err != nil {
		return 0
	}
	n = bw.WriteU24(uint32(v), e)
	return n
}

// WriteI32 ...
func (bw *Writer) WriteI32(v int32, e Endian) (n int) {
	if bw.err != nil {
		return 0
	}
	n = bw.WriteU32(uint32(v), e)
	return n
}

// WriteU8 ...
func (bw *Writer) WriteU8(v uint8) (n int) {
	if bw.err != nil {
		return 0
	}
	b := make([]byte, 1)
	b[0] = byte(v)
	n = bw.writeBytes(b)

	return n
}

// WriteU16 ...
func (bw *Writer) WriteU16(v uint16, e Endian) (n int) {
	if bw.err != nil {
		return 0
	}

	b := make([]byte, 2)

	if e == LittleEndian {
		b[0] = byte(v & 0xFF)
		b[1] = byte((v >> 8) & 0xFF)
	} else {
		b[0] = byte((v >> 8) & 0xFF)
		b[1] = byte(v & 0xFF)
	}

	n = bw.writeBytes(b)

	return n
}

// WriteU24 ...
func (bw *Writer) WriteU24(v uint32, e Endian) (n int) {
	if bw.err != nil {
		return 0
	}

	b := make([]byte, 3)

	if e == LittleEndian {
		b[0] = byte(v & 0xFF)
		b[1] = byte((v >> 8) & 0xFF)
		b[2] = byte((v >> 16) & 0xFF)
	} else {
		b[0] = byte((v >> 16) & 0xFF)
		b[1] = byte((v >> 8) & 0xFF)
		b[2] = byte(v & 0xFF)
	}

	n = bw.writeBytes(b)

	return n
}

// WriteU32 ...
func (bw *Writer) WriteU32(v uint32, e Endian) (n int) {
	if bw.err != nil {
		return 0
	}

	b := make([]byte, 4)

	if e == LittleEndian {
		b[0] = byte(v & 0xFF)
		b[1] = byte((v >> 8) & 0xFF)
		b[2] = byte((v >> 16) & 0xFF)
		b[3] = byte((v >> 24) & 0xFF)
	} else {
		b[0] = byte((v >> 24) & 0xFF)
		b[1] = byte((v >> 16) & 0xFF)
		b[2] = byte((v >> 8) & 0xFF)
		b[3] = byte(v & 0xFF)
	}

	n = bw.writeBytes(b)

	return n
}

// WriteS8 ...
func (bw *Writer) WriteS8(s string) (n int) {
	if bw.err != nil {
		return 0
	}

	var v uint8
	v = uint8(s[0])
	n = bw.WriteU8(v)

	return n
}

// WriteS16 ...
func (bw *Writer) WriteS16(s string, e Endian) (n int) {
	if bw.err != nil {
		return 0
	}
	var v uint16
	v = uint16(s[0])<<8 |
		uint16(s[1])
	n = bw.WriteU16(v, e)

	return n
}

// WriteS24 ...
func (bw *Writer) WriteS24(s string, e Endian) (n int) {
	if bw.err != nil {
		return 0
	}
	var v uint32
	v = uint32(s[0])<<16 |
		uint32(s[1])<<8 |
		uint32(s[2])
	n = bw.WriteU24(v, e)

	return n
}

// WriteS32 ...
func (bw *Writer) WriteS32(s string, e Endian) (n int) {
	if bw.err != nil {
		return 0
	}
	var v uint32
	v = uint32(s[0])<<24 |
		uint32(s[1])<<16 |
		uint32(s[2])<<8 |
		uint32(s[3])
	n = bw.WriteU32(v, e)

	return n
}
