// 9-2
package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.OpenFile("hello.txt.gz")

	if err != nil {
		fmt.Fprintln(err)
		return
	}

	defer file.Close()

	r, err := gzip.NewReader(file)

	if err != nil {
		fmt.Fprintln(err)
		return
	}

	defer r.Close()
	b, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Fprintln(err)
		return
	}

	fmt.Println(string(b))
}
