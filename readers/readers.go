package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (mr MyReader) Read(b []byte) (n int, err error) {
	for idx := range b {
		b[idx] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
