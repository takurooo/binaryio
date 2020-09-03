# binaryio
Binary reader/writer for Go.


# godoc
https://godoc.org/github.com/takurooo/binaryio

# Examples

```go
package main

import (
	"os"

	bio "github.com/takurooo/binaryio"
)

func main() {

	var v uint32
	v = 0x12345678

	// ----------------------
	// Writer Examples
	// ----------------------
	fw, _ := os.Create("temp.bin")
	w := bio.NewWriter(fw)

	w.WriteU32(v, bio.LittleEndian)
	if w.Err() != nil {
		panic(w.Err())
	}

	w.WriteU32(v, bio.BigEndian)
	if w.Err() != nil {
		panic(w.Err())
	}

	fw.Close()

	// ----------------------
	// Reader Examples
	// ----------------------
	fr, _ := os.Open("temp.bin")
	r := bio.NewReader(fr)

	if r.ReadU32(bio.LittleEndian) != v {
		panic(w.Err())
	}

	if r.ReadU32(bio.BigEndian) != v {
		panic(w.Err())
	}

	fr.Close()
}
