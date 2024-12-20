// 9-1
package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	
	file, err := os.OpenFile(
		"hello.txt.gz",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644)
	)
	
	if err != nil {
		fmt.Fprintln(err)
		return
	}
	
	defer file.Close()
	
	w := gzip.NewWriter(file)
	
	defer w.Close()
	
	s := "Hello, world!"
	w.Write([]byte(s))
	w.Flush()
}
