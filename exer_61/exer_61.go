package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13Calc(p byte) byte {
	if (p >= 'a' && p <= 'm') || (p >= 'A' && p <= 'M') {
		p += 13
	} else if (p >= 'n' && p <= 'z') || (p >= 'N' && p <= 'Z') {
		p -= 13
	}

	return p
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)

	for i := 0; i < n; i++ {
		p[i] = rot13Calc(p[i])
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
