package main

import (
	"fmt"
	"os"
)

func main()  {
	// 打开文件
	filename := "/Users/aizsfgk/app"
	file, err := os.Open(filename)
	if err != nil {
		panic("os.Open err: "+err.Error())
	}

	fileInfos, err := file.Readdir(-1)
	if err != nil {
		panic("file.Readdir err: "+err.Error())
	}

	fmt.Printf("fileInfos: %#v\n", fileInfos)



	//file, err := os.Open(filename)
	//if err != nil {
	//	panic("os.Open err: "+err.Error())
	//}
	//
	//fileInfos, err := file.Readdir(0)
	//if err != nil {
	//	panic("file.Readdir err: "+err.Error())
	//}
	//
	//fmt.Printf("fileInfos: %#v\n", fileInfos)



	return
}
