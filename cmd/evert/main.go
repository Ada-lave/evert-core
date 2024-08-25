package main

import (
	"fmt"
	"os"

	"github.com/Ada-lave/evert-core"
)

func main() {
	formatDoc()
}

func formatDoc() {
	file, err := os.Open("test_docs/origin.docx")

	if err != nil {
		panic(err)
	}

	fileinfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	evertDoc, err := evert.New(file, fileinfo.Size())

	if err != nil {
		panic(err)
	}

	evertFormatter := evert.NewFormatter(evertDoc)
	evertFormatter.Format(evert.FormatterParams{
		AddSpacesBeetweenImageText: true,
		FormatImagDescription: false,
	})
	
	res, err := evertDoc.GetBytes()

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
