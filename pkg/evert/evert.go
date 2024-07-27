package evert

import (
	"io"
	"os"

	"github.com/fumiama/go-docx"
)

type EvertDoc struct{
	Doc *docx.Docx
}

func (ED *EvertDoc) FormatImageCaption() {
	for idx, el := range ED.Doc.Document.Body.Items {
		switch element := el.(type) {
		case *docx.Paragraph:
			for _, paragraphChildren := range element.Children {
				switch paragraphElement := paragraphChildren.(type) {
				case *docx.Run:
					if ED.checkHaveDrawing(paragraphElement) && ED.IsHaveEmptySpace(idx + 1, &ED.Doc.Document.Body.Items) {
						
						ED.AddSpace(idx + 1, &ED.Doc.Document.Body.Items)
					}
				}
			} 
		}
	}

	ED.SaveFormattedDoc()
}

func (ED *EvertDoc) SaveFormattedDoc() {
	f, err := os.Create("generated.docx")

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

func(ED *EvertDoc) AddSpace(idx int, elements *[]interface{}) {
	buff := make([]interface{}, len((*elements)[idx+1:]))
	copy(buff, (*elements)[idx+1:])
	*elements = (*elements)[:idx+1]
	*elements = append(*elements, ED.Doc.AddParagraph().AddText("\n"))
	*elements = append(*elements, buff...)
}

func (ED *EvertDoc) IsHaveEmptySpace(idx int, elements *[]interface{}) bool {
	for _, el := range (*elements)[idx].(*docx.Paragraph).Children {
		switch run := el.(type) {
		case *docx.Run:
			for _, runElement := range run.Children {
				switch text := runElement.(type) {
				case *docx.Text:
					if text.Text != "\n" {
						return true
					}
				}
			} 
		}
	}

	return false
}

func (ED *EvertDoc) checkHaveDrawing(elements *docx.Run) bool {
	for _, runChildren := range elements.Children {
		switch runChildren.(type) {
		case *docx.Drawing:
			return true
		}
	}

	return false
}

func New(file io.ReaderAt, size int64) (*EvertDoc, error) {
	doc, err := docx.Parse(file, size)

	if err != nil {
		return &EvertDoc{}, err
	}

	return &EvertDoc{Doc: doc}, nil
} 
