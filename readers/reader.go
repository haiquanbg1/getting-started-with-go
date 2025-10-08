package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (MyReader) Read(a []byte) (int, error) {
	for i := range a {
		a[i] = 'A'
	}

	return len(a), nil
}

func main() {
	reader.Validate(MyReader{})
}
