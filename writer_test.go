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
		// WriteI8
		var n int
		v := int8(math.MinInt8)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteI8(v)
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
		if r.ReadI8() != v {
			t.Fatalf("Invalid ReadU8 %x", v)
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
}
