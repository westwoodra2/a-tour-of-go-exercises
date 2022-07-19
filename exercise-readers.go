package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (m MyReader) Read(b []byte) (int, error) {
	copy(b, []byte{'A'})
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
