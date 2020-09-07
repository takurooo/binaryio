package binaryio

import (
	"io"
)

// Reader ...
type Reader struct {
	io.ReaderAt
	offset int64
	err    error
	sbuf64 []byte
	b64    []byte
}

// NewReader ...
func NewReader(r io.ReaderAt) (br *Reader) {
	br = &Reader{
		r,
		0,
		nil,
		make([]byte, 8),
		make([]byte, 8),
	}
	return br
}

func (br *Reader) readBytes(n uint64) []byte {
	data := br.b64[:n]
	_, br.err = br.ReadAt(data, br.offset)
	br.offset += int64(n)
	return data
}

func (br *Reader) setErr(err error) {
	br.err = err
}

// Err ...
func (br *Reader) Err() error {
	return br.err
}

// ReadRaw ...
func (br *Reader) ReadRaw(n uint64) []byte {
	if br.err != nil {
		return nil
	}
	return br.readBytes(n)
}

// ReadI8 ...
func (br *Reader) ReadI8() int8 {
	if br.err != nil {
		return 0
	}
	return int8(br.ReadU8())
}

// ReadI16 ...
func (br *Reader) ReadI16(e Endian) int16 {
	if br.err != nil {
		return 0
	}
	return int16(br.ReadU16(e))
}

// ReadI24 ...
func (br *Reader) ReadI24(e Endian) int32 {
	if br.err != nil {
		return 0
	}
	tmp := br.ReadU24(e)

	var b int32
	if tmp&0x800000 == 0x800000 {
		b = -8388608 + int32(tmp&0x7FFFFF)
	} else {
		b = int32(tmp)
	}

	return b
}

// ReadI32 ...
func (br *Reader) ReadI32(e Endian) int32 {
	if br.err != nil {
		return 0
	}
	return int32(br.ReadU32(e))
}

// ReadI64 ...
func (br *Reader) ReadI64(e Endian) int64 {
	if br.err != nil {
		return 0
	}
	return int64(br.ReadU64(e))
}

// ReadU8 ...
func (br *Reader) ReadU8() uint8 {
	if br.err != nil {
		return 0
	}
	data := br.readBytes(1)
	return uint8(data[0])
}

// ReadU16 ...
func (br *Reader) ReadU16(e Endian) uint16 {
	if br.err != nil {
		return 0
	}
	data := br.readBytes(2)

	var b uint16
	if e == LittleEndian {
		b = uint16(data[1])<<8 +
			uint16(data[0])
	} else {
		b = uint16(data[0])<<8 +
			uint16(data[1])
	}

	return b
}

// ReadU24 ...
func (br *Reader) ReadU24(e Endian) uint32 {
	if br.err != nil {
		return 0
	}
	data := br.readBytes(3)

	var b uint32
	if e == LittleEndian {
		b = uint32(data[2])<<16 +
			uint32(data[1])<<8 +
			uint32(data[0])
	} else {
		b = uint32(data[0])<<16 +
			uint32(data[1])<<8 +
			uint32(data[2])
	}

	return b
}

// ReadU32 ...
func (br *Reader) ReadU32(e Endian) uint32 {
	if br.err != nil {
		return 0
	}
	data := br.readBytes(4)

	var b uint32
	if e == LittleEndian {
		b = uint32(data[3])<<24 +
			uint32(data[2])<<16 +
			uint32(data[1])<<8 +
			uint32(data[0])
	} else {
		b = uint32(data[0])<<24 +
			uint32(data[1])<<16 +
			uint32(data[2])<<8 +
			uint32(data[3])
	}

	return b
}

// ReadU64 ...
func (br *Reader) ReadU64(e Endian) uint64 {
	if br.err != nil {
		return 0
	}
	data := br.readBytes(8)

	var b uint64
	if e == LittleEndian {
		b = uint64(data[7])<<56 +
			uint64(data[6])<<48 +
			uint64(data[5])<<40 +
			uint64(data[4])<<32 +
			uint64(data[3])<<24 +
			uint64(data[2])<<16 +
			uint64(data[1])<<8 +
			uint64(data[0])
	} else {
		b = uint64(data[0])<<56 +
			uint64(data[1])<<48 +
			uint64(data[2])<<40 +
			uint64(data[3])<<32 +
			uint64(data[4])<<24 +
			uint64(data[5])<<16 +
			uint64(data[6])<<8 +
			uint64(data[7])
	}

	return b
}

// ReadS8 ...
func (br *Reader) ReadS8() string {
	if br.err != nil {
		return ""
	}
	return string(br.ReadU8())
}

// ReadS16 ...
func (br *Reader) ReadS16(e Endian) string {
	if br.err != nil {
		return ""
	}
	data := br.ReadU16(e)

	br.sbuf64[1] = byte(data)
	br.sbuf64[0] = byte(data >> 8)

	return string(br.sbuf64[:2])
}

// ReadS24 ...
func (br *Reader) ReadS24(e Endian) string {
	if br.err != nil {
		return ""
	}
	data := br.ReadU24(e)

	br.sbuf64[2] = byte(data)
	br.sbuf64[1] = byte(data >> 8)
	br.sbuf64[0] = byte(data >> 16)

	return string(br.sbuf64[:3])
}

// ReadS32 ...
func (br *Reader) ReadS32(e Endian) string {
	if br.err != nil {
		return ""
	}
	data := br.ReadU32(e)

	br.sbuf64[3] = byte(data)
	br.sbuf64[2] = byte(data >> 8)
	br.sbuf64[1] = byte(data >> 16)
	br.sbuf64[0] = byte(data >> 24)

	return string(br.sbuf64[:4])
}

// ReadS64 ...
func (br *Reader) ReadS64(e Endian) string {
	if br.err != nil {
		return ""
	}
	data := br.ReadU64(e)

	br.sbuf64[7] = byte(data)
	br.sbuf64[6] = byte(data >> 8)
	br.sbuf64[5] = byte(data >> 16)
	br.sbuf64[4] = byte(data >> 24)
	br.sbuf64[3] = byte(data >> 32)
	br.sbuf64[2] = byte(data >> 40)
	br.sbuf64[1] = byte(data >> 48)
	br.sbuf64[0] = byte(data >> 56)

	return string(br.sbuf64[:8])
}
