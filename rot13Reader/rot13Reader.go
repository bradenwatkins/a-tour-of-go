package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rotReader rot13Reader) Read(b []byte) (n int, e error) {
	n, e = rotReader.r.Read(b)
	for i := 0; i < n; i++ {
		val := b[i]
		if 'a' <= val && val <= 'z' {
			b[i] = 'a' + (val-'a'+13)%26
		} else if 'A' <= val && val <= 'Z' {
			b[i] = 'A' + (val-'A'+13)%26
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
