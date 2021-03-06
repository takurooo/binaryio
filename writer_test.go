package binaryio

import (
	"fmt"
	"math"
	"os"
	"testing"
)

func openWriteFile(s string, t *testing.T) *os.File {
	fw, err := os.Create(s)
	if err != nil {
		t.Fatal(err)
	}
	return fw
}

func openReadFile(s string, t *testing.T) *os.File {
	fr, err := os.Open(s)
	if err != nil {
		t.Fatal(err)
	}
	return fr
}

func removeFile(s string, t *testing.T) {
	if err := os.Remove(s); err != nil {
		fmt.Println(err)
	}
}

func TestWriter(t *testing.T) {
	testFileName := "test.bin"

	{
		var n int
		i8 := int8(math.MinInt8)
		i16 := int16(math.MinInt16)
		i32 := int32(math.MinInt32)
		i64 := int64(math.MinInt64)
		u8 := uint8(math.MaxUint8)
		u16 := uint16(math.MaxUint16)
		u32 := uint32(math.MaxUint32)
		u64 := uint64(math.MaxUint64)
		si8 := []int8{1, 2, 3}
		si16 := []int16{1, 2, 3}
		si32 := []int32{1, 2, 3}
		si64 := []int64{1, 2, 3}
		su8 := []uint8{1, 2, 3}
		su16 := []uint16{1, 2, 3}
		su32 := []uint32{1, 2, 3}
		su64 := []uint64{1, 2, 3}

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteX(
			LittleEndian,
			i8, i16, i32, i64,
			u8, u16, u32, u64,
			si8, si16, si32, si64,
			su8, su16, su32, su64)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 120 {
			t.Fatalf("Invalid WriteU8 %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadI8() != i8 {
			t.Fatalf("Invalid ReadI8")
		}
		if r.ReadI16(LittleEndian) != i16 {
			t.Fatalf("Invalid ReadI16")
		}
		if r.ReadI32(LittleEndian) != i32 {
			t.Fatalf("Invalid ReadI32")
		}
		if r.ReadI64(LittleEndian) != i64 {
			t.Fatalf("Invalid ReadI64")
		}
		if r.ReadU8() != u8 {
			t.Fatalf("Invalid ReadU8")
		}
		if r.ReadU16(LittleEndian) != u16 {
			t.Fatalf("Invalid ReadU16")
		}
		if r.ReadU32(LittleEndian) != u32 {
			t.Fatalf("Invalid ReadU32")
		}
		if r.ReadU64(LittleEndian) != u64 {
			t.Fatalf("Invalid ReadU64")
		}
		for _, v := range si8 {
			if r.ReadI8() != v {
				t.Fatalf("Invalid ReadI8")
			}
		}
		for _, v := range si16 {
			if r.ReadI16(LittleEndian) != v {
				t.Fatalf("Invalid ReadI16")
			}
		}
		for _, v := range si32 {
			if r.ReadI32(LittleEndian) != v {
				t.Fatalf("Invalid ReadI32")
			}
		}
		for _, v := range si64 {
			if r.ReadI64(LittleEndian) != v {
				t.Fatalf("Invalid ReadI64")
			}
		}
		for _, v := range su8 {
			if r.ReadU8() != v {
				t.Fatalf("Invalid ReadU8")
			}
		}
		for _, v := range su16 {
			if r.ReadU16(LittleEndian) != v {
				t.Fatalf("Invalid ReadU16")
			}
		}
		for _, v := range su32 {
			if r.ReadU32(LittleEndian) != v {
				t.Fatalf("Invalid ReadU32")
			}
		}
		for _, v := range su64 {
			if r.ReadU64(LittleEndian) != v {
				t.Fatalf("Invalid ReadU64")
			}
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteI16LE WriteI16BE
		var n int
		v := int16(math.MinInt16)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteI16(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 2 {
			t.Fatalf("Invalid WriteI16LE %d", n)
		}
		n = w.WriteI16(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 2 {
			t.Fatalf("Invalid WriteI16BE %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadI16(LittleEndian) != v {
			t.Fatalf("Invalid ReadI16LE %x", v)
		}
		if r.ReadI16(BigEndian) != v {
			t.Fatalf("Invalid ReadI16LE %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteI24LE WriteBEI24
		var n int
		v := int32(-8388608)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteI24(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 3 {
			t.Fatalf("Invalid WriteI24LE %d", n)
		}
		n = w.WriteI24(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 3 {
			t.Fatalf("Invalid WriteI24LE %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadI24(LittleEndian) != v {
			t.Fatalf("Invalid ReadI24LE %x", v)
		}
		if r.ReadI24(BigEndian) != v {
			t.Fatalf("Invalid ReadI24LE %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteI32LE WriteBEI32
		var n int
		v := int32(math.MinInt32)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteI32(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 4 {
			t.Fatalf("Invalid WriteI32LE %d", n)
		}
		n = w.WriteI32(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 4 {
			t.Fatalf("Invalid WriteI32LE %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadI32(LittleEndian) != v {
			t.Fatalf("Invalid ReadI32LE %x", v)
		}
		if r.ReadI32(BigEndian) != v {
			t.Fatalf("Invalid ReadI32LE %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteI64LE WriteI64BE
		var n int
		v := int64(math.MinInt64)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteI64(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 8 {
			t.Fatalf("Invalid WriteI64LE %d", n)
		}
		n = w.WriteI64(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 8 {
			t.Fatalf("Invalid WriteI64LE %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadI64(LittleEndian) != v {
			t.Fatalf("Invalid ReadI64LE %x", v)
		}
		if r.ReadI64(BigEndian) != v {
			t.Fatalf("Invalid ReadI64LE %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteU8
		var n int
		v := uint8(0x12)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteU8(v)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 1 {
			t.Fatalf("Invalid WriteU8 %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadU8() != v {
			t.Fatalf("Invalid ReadU8 %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteU16LE WriteU16BE
		var n int
		v := uint16(0x1234)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteU16(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 2 {
			t.Fatalf("Invalid WriteU16LE %d", n)
		}
		n = w.WriteU16(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 2 {
			t.Fatalf("Invalid WriteU16BE %d", n)
		}
		fw.Sync()
		fw.Close()

		var rv uint16
		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		rv = r.ReadU16(LittleEndian)
		if rv != v {
			t.Fatalf("Invalid ReadU16LE %x", rv)
		}
		rv = r.ReadU16(BigEndian)
		if rv != v {
			t.Fatalf("Invalid ReadU16BE %x", rv)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteU24LE WriteU24BE
		var n int
		v := uint32(0x123456)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteU24(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 3 {
			t.Fatalf("Invalid WriteU24LE %d", n)
		}
		n = w.WriteU24(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 3 {
			t.Fatalf("Invalid WriteU24BE %d", n)
		}
		fw.Sync()
		fw.Close()

		var rv uint32
		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		rv = r.ReadU24(LittleEndian)
		if rv != v {
			t.Fatalf("Invalid ReadU24LE %x", rv)
		}
		rv = r.ReadU24(BigEndian)
		if rv != v {
			t.Fatalf("Invalid ReadU24BE %x", rv)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteU32LE WriteU32BE
		var n int
		v := uint32(0x12345678)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteU32(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 4 {
			t.Fatalf("Invalid WriteU32LE %d", n)
		}
		n = w.WriteU32(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 4 {
			t.Fatalf("Invalid WriteU32BE %d", n)
		}
		fw.Sync()
		fw.Close()

		var rv uint32
		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		rv = r.ReadU32(LittleEndian)
		if rv != v {
			t.Fatalf("Invalid ReadU32LE %x", rv)
		}
		rv = r.ReadU32(BigEndian)
		if rv != v {
			t.Fatalf("Invalid ReadU32BE %x", rv)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteU64LE WriteU64BE
		var n int
		v := uint64(0x0123456789ABCDEF)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteU64(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 8 {
			t.Fatalf("Invalid WriteU64LE %d", n)
		}
		n = w.WriteU64(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 8 {
			t.Fatalf("Invalid WriteU64BE %d", n)
		}
		fw.Sync()
		fw.Close()

		var rv uint64
		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		rv = r.ReadU64(LittleEndian)
		if rv != v {
			t.Fatalf("Invalid ReadU64LE %x", rv)
		}
		rv = r.ReadU64(BigEndian)
		if rv != v {
			t.Fatalf("Invalid ReadU64BE %x", rv)
		}
		fr.Close()

		removeFile(testFileName, t)
	}

	{
		// WriteS8
		var n int
		v := "1"

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteS8(v)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 1 {
			t.Fatalf("Invalid WriteU8 %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadS8() != v {
			t.Fatalf("Invalid ReadU8 %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteS16LE WriteS16BE
		var n int
		v := "12"

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteS16(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 2 {
			t.Fatalf("Invalid WriteS16LE %d", n)
		}
		n = w.WriteS16(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 2 {
			t.Fatalf("Invalid WriteS16BE %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadS16(LittleEndian) != v {
			t.Fatalf("Invalid ReadS16LE %x", v)
		}
		if r.ReadS16(BigEndian) != v {
			t.Fatalf("Invalid ReadS16BE %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteS24LE WriteS24BE
		var n int
		v := "123"

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteS24(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 3 {
			t.Fatalf("Invalid WriteS24LE %d", n)
		}
		n = w.WriteS24(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 3 {
			t.Fatalf("Invalid WriteS24BE %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadS24(LittleEndian) != v {
			t.Fatalf("Invalid ReadS24LE %x", v)
		}
		if r.ReadS24(BigEndian) != v {
			t.Fatalf("Invalid ReadS24BE %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteS32LE WriteBES32
		var n int
		v := "1234"

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteS32(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 4 {
			t.Fatalf("Invalid WriteS32LE %d", n)
		}
		n = w.WriteS32(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 4 {
			t.Fatalf("Invalid WriteS32LE %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadS32(LittleEndian) != v {
			t.Fatalf("Invalid ReadS32LE %x", v)
		}
		if r.ReadS32(BigEndian) != v {
			t.Fatalf("Invalid ReadS32BE %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteS64LE WriteS64BE
		var n int
		v := "12345678"

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteS64(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 8 {
			t.Fatalf("Invalid WriteS64LE %d", n)
		}
		n = w.WriteS64(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 8 {
			t.Fatalf("Invalid WriteS64LE %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadS64(LittleEndian) != v {
			t.Fatalf("Invalid ReadS64LE %x", v)
		}
		if r.ReadS64(BigEndian) != v {
			t.Fatalf("Invalid ReadS64BE %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
}
