package evert

import (
	"io"
	"os"

	"github.com/fumiama/go-docx"
)

type Evert struct {
	Doc *docx.Docx
}

func (E *Evert) SaveFormattedDoc(path string) {
	f, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	_, err = E.Doc.WriteTo(f)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}

func (E *Evert) GetBytes() {
	
}

func New(file io.ReaderAt, size int64) (*Evert, error) {
	doc, err := docx.Parse(file, size)
	if err != nil {
		return &Evert{}, err
	}

	return &Evert{Doc: doc}, nil
}
