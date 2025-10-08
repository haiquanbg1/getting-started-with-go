package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return 'A' + (b-'A'+13)%26
	} else {
		if b >= 'a' && b <= 'z' {
			return 'a' + (b-'a'+13)%26
		}
	}

	return b
}

func (r *rot13Reader) Read(p []byte) (int, error) {
	n, error := r.r.Read(p)

	for i := 0; i < n; i++ {
		p[i] = rot13(p[i])
	}

	return n, error
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
