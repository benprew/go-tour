package main

import (
	"fmt"
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
		if n >= 'A' && n <= 'z' {
			fmt.Printf("%d %d", n, 'A')
			out[i] = (n+13)%26 + (n - 'A')
		} else {
			out[i] = n
		}

	}
	return l, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
