package main

import (
	"github.com/Ada-lave/evert-core/pkg/evert"
	"os"
)


func main() {

	file, err := os.Open("test.docx")

	if err != nil {
		panic(err)
	}

	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	evertDoc, err := evert.New(file, fileinfo.Size())
	evertDoc.FormatImageBody(false)
}