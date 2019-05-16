package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (mr MyReader) Read(b []byte) (n int, err error) {
	for idx, _ := range b {
		b[idx] = 'A'
	}
	return len(b), nil
}


// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
