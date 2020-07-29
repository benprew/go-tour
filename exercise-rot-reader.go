package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(out []byte) (int, error) {
	l, err := r.r.Read(out)
	if err != nil {
		return l, err
	}
	for i, n := range out {
		out[i] = n
		if n >= 'A' && n <= 'Z' {
			out[i] = rot13(n, 'A')
		} else if n >= 'a' && n < 'z' {
			out[i] = rot13(n, 'a')
		}
	}
	return l, nil
}

func rot13(val byte, offset byte) byte {
	return ((val+13)-offset)%26 + offset
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
