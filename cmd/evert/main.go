package main

import (
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
	evertFormatter.Format(true, true)
	evertDoc.SaveFormattedDoc("test_docs/generated.docx")
}
