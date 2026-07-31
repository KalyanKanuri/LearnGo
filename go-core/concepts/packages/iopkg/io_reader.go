package iopkg

import (
	"fmt"
	"io"
)

func IOReader(source io.Reader) {
	buf := make([]byte, 4*1024)
	bytesRead, err := source.Read(buf)
	if err != nil {
		if err == io.EOF {
			fmt.Println(string(buf[:bytesRead]))
			return
		}
		fmt.Println("Error reading bytes", err)
	}
	fmt.Println("io.Read()", string(buf[:bytesRead]))
}
