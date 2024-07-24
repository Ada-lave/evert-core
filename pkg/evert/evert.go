package evert

import (
	"fmt"
	"io"

	"github.com/fumiama/go-docx"
)

type EvertDoc struct{
	Doc *docx.Docx
}

func (ED *EvertDoc) FormatImageCaption() {
	for _, el := range ED.Doc.Document.Body.Items {
		fmt.Printf("%T\n", el)
	}

}

func New(file io.ReaderAt, size int64) (*EvertDoc, error) {
	doc, err := docx.Parse(file, size)

	if err != nil {
		return &EvertDoc{}, err
	}

	return &EvertDoc{Doc: doc}, nil
} 
