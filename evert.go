package evert

import (
	"bytes"
	"io"
	"os"
	"github.com/fumiama/go-docx"
)

type Evert struct {
	Doc *docx.Docx
}

func (E *Evert) SaveFormattedDoc(path string) error {
	f, err := os.Create(path)

	if err != nil {
		return err
	}
	_, err = E.Doc.WriteTo(f)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

func (E *Evert) GetBytes() ([]byte, error) {
	var buf bytes.Buffer

	_, err := E.Doc.WriteTo(&buf)
	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func New(file io.ReaderAt, size int64) (*Evert, error) {
	doc, err := docx.Parse(file, size)
	if err != nil {
		return &Evert{}, err
	}

	return &Evert{Doc: doc}, nil
}
