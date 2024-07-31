package main

import (
	"os"

	"github.com/Ada-lave/evert-core/pkg/evert"
	"github.com/Ada-lave/evert-core/pkg/formatter"
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
	evertFormatter := formatter.NewFormatter(evertDoc)
	evertFormatter.FormatImageBody(false)
	evertDoc.SaveFormattedDoc()
}