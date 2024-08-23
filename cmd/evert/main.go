package main

import (
	"os"

	"github.com/Ada-lave/evert-core/pkg/evert"
	"github.com/Ada-lave/evert-core/pkg/formatter"
)


func main() {
	formatDoc()
}

func formatDoc() {
	file, err := os.Open("origin.docx")

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

	evertFormatter := formatter.NewFormatter(evertDoc)
	evertFormatter.Format(true, true)
	evertDoc.SaveFormattedDoc()
}
