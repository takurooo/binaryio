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
		b := r.ReadI8()
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
		b := r.ReadI16(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != math.MinInt16 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x80}))
		b := r.ReadI24(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != -8388608 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x80}))
		b := r.ReadI32(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != math.MinInt32 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80}))
		b := r.ReadI64(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != math.MinInt64 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}

	// -----------------------------
	// ReadBEInt
	// -----------------------------
	{
		r := NewReader(bytes.NewReader([]byte{0x80, 0x00}))
		b := r.ReadI16(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != math.MinInt16 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x80, 0x00, 0x00}))
		b := r.ReadI24(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != -8388608 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x80, 0x00, 0x00, 0x00}))
		b := r.ReadI32(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != math.MinInt32 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}))
		b := r.ReadI64(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != math.MinInt64 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}

	// -----------------------------
	// ReadUint
	// -----------------------------
	{
		r := NewReader(bytes.NewReader([]byte{0x80}))
		b := r.ReadU8()
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
		b := r.ReadU16(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != 0x8000 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x80}))
		b := r.ReadU24(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != 0x800000 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x80}))
		b := r.ReadU32(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != 0x80000000 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80}))
		b := r.ReadU64(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != 0x8000000000000000 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	// -----------------------------
	// ReadBEUint
	// -----------------------------
	{
		r := NewReader(bytes.NewReader([]byte{0x80, 0x00}))
		b := r.ReadU16(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != 0x8000 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x80, 0x00, 0x00}))
		b := r.ReadU24(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != 0x800000 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x80, 0x00, 0x00, 0x00}))
		b := r.ReadU32(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != 0x80000000 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte{0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}))
		b := r.ReadU64(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != 0x8000000000000000 {
			t.Fatalf("Invalid Read Value: %d", b)
		}
	}
	// -----------------------------
	// ReadStr
	// -----------------------------
	{
		r := NewReader(bytes.NewReader([]byte("1")))
		b := r.ReadS8()
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
		b := r.ReadS16(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != "12" {
			t.Fatalf("Invalid Read Value: %s", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte("321")))
		b := r.ReadS24(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != "123" {
			t.Fatalf("Invalid Read Value: %s", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte("4321")))
		b := r.ReadS32(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != "1234" {
			t.Fatalf("Invalid Read Value: %s", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte("87654321")))
		b := r.ReadS64(LittleEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != "12345678" {
			t.Fatalf("Invalid Read Value: %s", b)
		}
	}
	// -----------------------------
	// ReadBEStr
	// -----------------------------
	{
		r := NewReader(bytes.NewReader([]byte("12")))
		b := r.ReadS16(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != "12" {
			t.Fatalf("Invalid Read Value: %s", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte("123")))
		b := r.ReadS24(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != "123" {
			t.Fatalf("Invalid Read Value: %s", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte("1234")))
		b := r.ReadS32(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != "1234" {
			t.Fatalf("Invalid Read Value: %s", b)
		}
	}
	{
		r := NewReader(bytes.NewReader([]byte("12345678")))
		b := r.ReadS64(BigEndian)
		if r.Err() != nil {
			t.Fatal(r.Err())
		}
		if b != "12345678" {
			t.Fatalf("Invalid Read Value: %s", b)
		}
	}
}
