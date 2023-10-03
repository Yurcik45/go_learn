package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Привіт, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}
	}
}

