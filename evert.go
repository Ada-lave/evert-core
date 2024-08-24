package evert

import (
	"io"
	"os"

	"github.com/fumiama/go-docx"
)

type EvertDoc struct {
	Doc *docx.Docx
}

func (ED *EvertDoc) SaveFormattedDoc(path string) {
	f, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	_, err = ED.Doc.WriteTo(f)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func New(file io.ReaderAt, size int64) (*EvertDoc, error) {
	doc, err := docx.Parse(file, size)
	if err != nil {
		return &EvertDoc{}, err
	}

	return &EvertDoc{Doc: doc}, nil
}
