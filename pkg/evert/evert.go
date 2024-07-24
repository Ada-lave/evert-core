package evert

import (
	"fmt"
	"io"
	"os"

	"github.com/fumiama/go-docx"
)

type EvertDoc struct{
	Doc *docx.Docx
}

func (ED *EvertDoc) FormatImageCaption() {
	for idx, el := range ED.Doc.Document.Body.Items {
		fmt.Printf("%T\n", el)
		switch element := el.(type) {
		case *docx.Paragraph:
			// fmt.Printf("%#v\n", element.Children...)
			for _, paragraphChildren := range element.Children {
				switch paragraphElement := paragraphChildren.(type) {
				case *docx.Run:
					if ED.checkHaveDrawing(paragraphElement) {
						// fmt.Printf("%T\n", paragraphElement.Children...)
						ED.AddSpace(idx, &ED.Doc.Document.Body.Items, el.(*docx.Paragraph))
						// paragraphElement.Children[idx+1] = append(paragraphElement.Children, element.AddText("\n")) 
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

func(ED *EvertDoc) AddSpace(idx int, elements *[]interface{}, parent *docx.Paragraph) {
	buff := (*elements)[idx:]
	*elements = (*elements)[:idx]
	(*elements)[idx] = parent.AddText("\n")
	*elements = append(*elements, buff...)

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
