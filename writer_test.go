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
		// WriteInt8
		var n int
		v := int8(math.MinInt8)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteInt8(v)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 1 {
			t.Fatalf("Invalid WriteUint8 %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadInt8() != v {
			t.Fatalf("Invalid ReadUint8 %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteLEInt16 WriteBEInt16
		var n int
		v := int16(math.MinInt16)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteInt16(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 2 {
			t.Fatalf("Invalid WriteLEInt16 %d", n)
		}
		n = w.WriteInt16(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 2 {
			t.Fatalf("Invalid WriteBEInt16 %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadInt16(LittleEndian) != v {
			t.Fatalf("Invalid ReadLEInt16 %x", v)
		}
		if r.ReadInt16(BigEndian) != v {
			t.Fatalf("Invalid ReadLEInt16 %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteLEInt24 WriteBEInt24
		var n int
		v := int32(-8388608)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteInt24(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 3 {
			t.Fatalf("Invalid WriteLEInt24 %d", n)
		}
		n = w.WriteInt24(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 3 {
			t.Fatalf("Invalid WriteLEInt24 %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadInt24(LittleEndian) != v {
			t.Fatalf("Invalid ReadLEInt24 %x", v)
		}
		if r.ReadInt24(BigEndian) != v {
			t.Fatalf("Invalid ReadLEInt24 %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteLEInt32 WriteBEInt32
		var n int
		v := int32(math.MinInt32)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteInt32(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 4 {
			t.Fatalf("Invalid WriteLEInt32 %d", n)
		}
		n = w.WriteInt32(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 4 {
			t.Fatalf("Invalid WriteLEInt32 %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadInt32(LittleEndian) != v {
			t.Fatalf("Invalid ReadLEInt32 %x", v)
		}
		if r.ReadInt32(BigEndian) != v {
			t.Fatalf("Invalid ReadLEInt32 %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}

	{
		// WriteUint8
		var n int
		v := uint8(0x12)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteUint8(v)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 1 {
			t.Fatalf("Invalid WriteUint8 %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadUint8() != v {
			t.Fatalf("Invalid ReadUint8 %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteLEUint16 WriteBEUint16
		var n int
		v := uint16(0x1234)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteUint16(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 2 {
			t.Fatalf("Invalid WriteLEUint16 %d", n)
		}
		n = w.WriteUint16(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 2 {
			t.Fatalf("Invalid WriteBEUint16 %d", n)
		}
		fw.Sync()
		fw.Close()

		var rv uint16
		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		rv = r.ReadUint16(LittleEndian)
		if rv != v {
			t.Fatalf("Invalid ReadLEUint16 %x", rv)
		}
		rv = r.ReadUint16(BigEndian)
		if rv != v {
			t.Fatalf("Invalid ReadBEUint16 %x", rv)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteLEUint24 WriteBEUint24
		var n int
		v := uint32(0x123456)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteUint24(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 3 {
			t.Fatalf("Invalid WriteLEUint24 %d", n)
		}
		n = w.WriteUint24(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 3 {
			t.Fatalf("Invalid WriteBEUint24 %d", n)
		}
		fw.Sync()
		fw.Close()

		var rv uint32
		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		rv = r.ReadUint24(LittleEndian)
		if rv != v {
			t.Fatalf("Invalid ReadLEUint24 %x", rv)
		}
		rv = r.ReadUint24(BigEndian)
		if rv != v {
			t.Fatalf("Invalid ReadBEUint24 %x", rv)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteLEUint32 WriteBEUint32
		var n int
		v := uint32(0x12345678)

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteUint32(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 4 {
			t.Fatalf("Invalid WriteLEUint32 %d", n)
		}
		n = w.WriteUint32(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 4 {
			t.Fatalf("Invalid WriteBEUint32 %d", n)
		}
		fw.Sync()
		fw.Close()

		var rv uint32
		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		rv = r.ReadUint32(LittleEndian)
		if rv != v {
			t.Fatalf("Invalid ReadLEUint32 %x", rv)
		}
		rv = r.ReadUint32(BigEndian)
		if rv != v {
			t.Fatalf("Invalid ReadBEUint32 %x", rv)
		}
		fr.Close()

		removeFile(testFileName, t)
	}

	{
		// WriteStr8
		var n int
		v := "1"

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteStr8(v)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 1 {
			t.Fatalf("Invalid WriteUint8 %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadStr8() != v {
			t.Fatalf("Invalid ReadUint8 %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteLEStr16 WriteBEStr16
		var n int
		v := "12"

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteStr16(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 2 {
			t.Fatalf("Invalid WriteLEStr16 %d", n)
		}
		n = w.WriteStr16(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 2 {
			t.Fatalf("Invalid WriteBEStr16 %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadStr16(LittleEndian) != v {
			t.Fatalf("Invalid ReadLEStr16 %x", v)
		}
		if r.ReadStr16(BigEndian) != v {
			t.Fatalf("Invalid ReadBEStr16 %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteLEStr24 WriteBEStr24
		var n int
		v := "123"

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteStr24(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 3 {
			t.Fatalf("Invalid WriteLEStr24 %d", n)
		}
		n = w.WriteStr24(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 3 {
			t.Fatalf("Invalid WriteBEStr24 %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadStr24(LittleEndian) != v {
			t.Fatalf("Invalid ReadLEStr24 %x", v)
		}
		if r.ReadStr24(BigEndian) != v {
			t.Fatalf("Invalid ReadBEStr24 %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
	{
		// WriteLEStr32 WriteBEStr32
		var n int
		v := "1234"

		fw := openWriteFile(testFileName, t)
		w := NewWriter(fw)
		n = w.WriteStr32(v, LittleEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 4 {
			t.Fatalf("Invalid WriteLEStr32 %d", n)
		}
		n = w.WriteStr32(v, BigEndian)
		if w.Err() != nil {
			t.Fatal(w.Err())
		}
		if n != 4 {
			t.Fatalf("Invalid WriteBEStr32 %d", n)
		}
		fw.Sync()
		fw.Close()

		fr := openReadFile(testFileName, t)
		r := NewReader(fr)
		if r.ReadStr32(LittleEndian) != v {
			t.Fatalf("Invalid ReadLEStr32 %x", v)
		}
		if r.ReadStr32(BigEndian) != v {
			t.Fatalf("Invalid ReadBEStr32 %x", v)
		}
		fr.Close()

		removeFile(testFileName, t)
	}
}
