package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		新建一个文件
	 */
	filename := "./file.txt"
	file, err := os.Create(filename)
	if err != nil {
		panic("os.Open err: "+err.Error())
	}

	data := []byte("create a new file, to write some data")
	writeN, err := file.Write(data)
	if err != nil {
		panic("file.Write err: "+err.Error())
	}
	fmt.Printf("\nfile writeN: %d ; data-length: %d\n", writeN, len(data))

	return
}
