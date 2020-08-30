package binaryio

import (
	"io"
)

// Reader ...
type Reader struct {
	io.ReaderAt
	offset int64
	err    error
}

// NewReader ...
func NewReader(r io.ReaderAt) (br *Reader) {
	br = &Reader{r, 0, nil}
	return br
}

func (br *Reader) readBytes(n uint64) []byte {
	offset := br.offset
	data := make([]byte, n)

	_, err := br.ReadAt(data, offset)
	br.setErr(err)

	defer func() {
		br.offset += int64(n)
	}()

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
		return []byte{}
	}
	data := br.readBytes(n)

	return data
}

// ReadInt8 ...
func (br *Reader) ReadInt8() int8 {
	if br.err != nil {
		return 0
	}

	return int8(br.ReadUint8())
}

// ReadInt16 ...
func (br *Reader) ReadInt16(e Endian) int16 {
	if br.err != nil {
		return 0
	}

	return int16(br.ReadUint16(e))
}

// ReadInt24 ...
func (br *Reader) ReadInt24(e Endian) int32 {
	if br.err != nil {
		return 0
	}
	tmp := br.ReadUint24(e)

	var b int32
	if tmp&0x800000 == 0x800000 {
		b = -8388608 + int32(tmp&0x7FFFFF)
	} else {
		b = int32(tmp)
	}

	return b
}

// ReadInt32 ...
func (br *Reader) ReadInt32(e Endian) int32 {
	if br.err != nil {
		return 0
	}

	return int32(br.ReadUint32(e))
}

// ReadUint8 ...
func (br *Reader) ReadUint8() uint8 {
	if br.err != nil {
		return 0
	}
	data := br.readBytes(1)

	return uint8(data[0])
}

// ReadUint16 ...
func (br *Reader) ReadUint16(e Endian) uint16 {
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

// ReadUint24 ...
func (br *Reader) ReadUint24(e Endian) uint32 {
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

// ReadUint32 ...
func (br *Reader) ReadUint32(e Endian) uint32 {
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

// ReadStr8 ...
func (br *Reader) ReadStr8() string {
	if br.err != nil {
		return ""
	}

	return string(br.ReadUint8())
}

// ReadStr16 ...
func (br *Reader) ReadStr16(e Endian) string {
	if br.err != nil {
		return ""
	}
	data := br.ReadUint16(e)

	tmp := make([]byte, 2)
	tmp[1] = byte(0x000000FF & data)
	tmp[0] = byte(0x000000FF & (data >> 8))

	return string(tmp)
}

// ReadStr24 ...
func (br *Reader) ReadStr24(e Endian) string {
	if br.err != nil {
		return ""
	}
	data := br.ReadUint24(e)

	tmp := make([]byte, 3)
	tmp[2] = byte(0x000000FF & data)
	tmp[1] = byte(0x000000FF & (data >> 8))
	tmp[0] = byte(0x000000FF & (data >> 16))

	return string(tmp)
}

// ReadStr32 ...
func (br *Reader) ReadStr32(e Endian) string {
	if br.err != nil {
		return ""
	}
	data := br.ReadUint32(e)

	tmp := make([]byte, 4)
	tmp[3] = byte(0x000000FF & data)
	tmp[2] = byte(0x000000FF & (data >> 8))
	tmp[1] = byte(0x000000FF & (data >> 16))
	tmp[0] = byte(0x000000FF & (data >> 24))

	return string(tmp)
}
