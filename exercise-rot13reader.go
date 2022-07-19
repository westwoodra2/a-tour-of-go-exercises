package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(x byte) byte {
	upperCase := x >= 'A' && x <= 'Z'
	lowerCase := x >= 'a' && x <= 'z'

	if !upperCase && !lowerCase {
		return x
	}

	x += 13
	if upperCase && x > 'Z' || lowerCase && x > 'z' {
		x -= 26
	}

	return x
}

func (rot rot13Reader) Read(buffer []byte) (int, error) {
	tmp := make([]byte, len(buffer))
	n, err := rot.r.Read(tmp)

	if err != nil {
		return n, err
	}

	for i, val := range tmp {
		buffer[i] = rot13(val)
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
