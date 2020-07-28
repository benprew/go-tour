package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (MyReader) Read(out []byte) (int, error) {
	for i := 0; i < len(out); i++ {
		out[i] = 'A'
	}

	return len(out), nil
}

func main() {
	reader.Validate(MyReader{})
}
