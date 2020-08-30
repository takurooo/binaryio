# binaryio
Binary read/write library for Go

# Usage

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

	w.WriteUint32(v, bio.LittleEndian)
	if w.Err() != nil {
		panic(w.Err())
	}

	w.WriteUint32(v, bio.BigEndian)
	if w.Err() != nil {
		panic(w.Err())
	}

	fw.Close()

	// ----------------------
	// Reader Examples
	// ----------------------
	fr, _ := os.Open("temp.bin")
	r := bio.NewReader(fr)

	if r.ReadUint32(bio.LittleEndian) != v {
		panic(w.Err())
	}

	if r.ReadUint32(bio.BigEndian) != v {
		panic(w.Err())
	}

	fr.Close()
}
```
