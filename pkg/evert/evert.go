package evert

import (
	"io"
	"os"

	"github.com/fumiama/go-docx"
)

type EvertDoc struct{
	Doc *docx.Docx
}

func (ED *EvertDoc) FormatImageBody(formatImageDescription bool) {
	itemsLen := len(ED.Doc.Document.Body.Items)
	for idx := 0; idx < itemsLen; idx++ {
		switch element := ED.Doc.Document.Body.Items[idx].(type) {
		case *docx.Paragraph:
			for _, paragraphChildren := range element.Children {
				switch paragraphElement := paragraphChildren.(type) {
				case *docx.Run:
					if ED.checkHaveDrawing(paragraphElement) && ED.IsHaveEmptySpace(idx + 1, &ED.Doc.Document.Body.Items) {
						ED.AddSpace(idx + 1, &ED.Doc.Document.Body.Items)
						itemsLen++
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
	buff := make([]interface{}, len((ED.Doc.Document.Body.Items)[idx+1:]))
	copy(buff, (ED.Doc.Document.Body.Items)[idx+1:])
	ED.Doc.Document.Body.Items = (ED.Doc.Document.Body.Items)[:idx+1]
	ED.Doc.Document.Body.Items = append(ED.Doc.Document.Body.Items, ED.Doc.AddParagraph().AddText("\n"))
	ED.Doc.Document.Body.Items = append(ED.Doc.Document.Body.Items, buff...)
}

func (ED *EvertDoc) IsHaveEmptySpace(idx int, elements *[]interface{}) bool {
	for _, el := range (ED.Doc.Document.Body.Items)[idx].(*docx.Paragraph).Children {
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

func (ED *EvertDoc) CapitalizeSentence(idx int) {

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
