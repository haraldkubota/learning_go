package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(c byte) byte {
	switch {
		case c>='A' && c <='M' || c>='a' && c <='m': return c+13
		case c >'M' && c <='Z' || c > 'm' && c <= 'z': return c-13
	}
	return c
}

func (a rot13Reader) Read(b []byte) (int, error) {
	n, err := a.r.Read(b)
	for i:=0; i<n; i+=1 {
		c:=b[i]
		b[i]=rot13(c)
	}
		
	return n, err
}


func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
