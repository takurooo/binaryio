package binaryio

import (
	"bytes"
	"math"
	"testing"
)

func TestReader(t *testing.T) {

	// -----------------------------
	// ReadInt
	// -----------------------------
	{
		r := NewReader(bytes.NewReader([]byte{0x80}))
		b := r.ReadInt8()
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != math.MinInt8 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	// -----------------------------
	// ReadLEInt
	// -----------------------------
	{
		r := NewReader(bytes.NewReader([]byte{0x00, 0x80}))
		b := r.ReadInt16(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != math.MinInt16 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x80}))
		b := r.ReadInt24(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != -8388608 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x80}))
		b := r.ReadInt32(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != math.MinInt32 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}

	// -----------------------------
	// ReadBEInt
	// -----------------------------
	{
		r := NewReader(bytes.NewReader([]byte{0x80, 0x00}))
		b := r.ReadInt16(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != math.MinInt16 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x80, 0x00, 0x00}))
		b := r.ReadInt24(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != -8388608 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x80, 0x00, 0x00, 0x00}))
		b := r.ReadInt32(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != math.MinInt32 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}

	// -----------------------------
	// ReadUint
	// -----------------------------
	{
		r := NewReader(bytes.NewReader([]byte{0x80}))
		b := r.ReadUint8()
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != 0x80 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	// -----------------------------
	// ReadLEUint
	// -----------------------------
	{
		r := NewReader(bytes.NewReader([]byte{0x00, 0x80}))
		b := r.ReadUint16(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != 0x8000 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x80}))
		b := r.ReadUint24(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != 0x800000 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x80}))
		b := r.ReadUint32(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != 0x80000000 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	// -----------------------------
	// ReadBEUint
	// -----------------------------
	{
		r := NewReader(bytes.NewReader([]byte{0x80, 0x00}))
		b := r.ReadUint16(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != 0x8000 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x80, 0x00, 0x00}))
		b := r.ReadUint24(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != 0x800000 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x80, 0x00, 0x00, 0x00}))
		b := r.ReadUint32(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != 0x80000000 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	// -----------------------------
	// ReadStr
	// -----------------------------
	{
		r := NewReader(bytes.NewReader([]byte("1")))
		b := r.ReadStr8()
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != "1" {
			t.Fatalf("Invalid Read Value: %s", b)
		}
	}
	// -----------------------------
	// ReadLEStr
	// -----------------------------
	{
		r := NewReader(bytes.NewReader([]byte("21")))
		b := r.ReadStr16(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != "12" {
			t.Fatalf("Invalid Read Value: %s", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte("321")))
		b := r.ReadStr24(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != "123" {
			t.Fatalf("Invalid Read Value: %s", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte("4321")))
		b := r.ReadStr32(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != "1234" {
			t.Fatalf("Invalid Read Value: %s", b)
		}
	}
	// -----------------------------
	// ReadBEStr
	// -----------------------------
	{
		r := NewReader(bytes.NewReader([]byte("12")))
		b := r.ReadStr16(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != "12" {
			t.Fatalf("Invalid Read Value: %s", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte("123")))
		b := r.ReadStr24(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != "123" {
			t.Fatalf("Invalid Read Value: %s", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte("1234")))
		b := r.ReadStr32(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != "1234" {
			t.Fatalf("Invalid Read Value: %s", b)
		}
	}
}
