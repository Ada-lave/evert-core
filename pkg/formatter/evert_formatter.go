package formatter

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/Ada-lave/evert-core/pkg/evert"
	"github.com/fumiama/go-docx"
)

type EvertFormatter struct {
	EvertDoc *evert.EvertDoc
}

func (EF *EvertFormatter) Format(addSpacesBeetweenImageText bool, formatImagDescription bool) {
	itemsLen := len(EF.EvertDoc.Doc.Document.Body.Items)
	for idx := 0; idx < itemsLen; idx++ {
		switch element := EF.EvertDoc.Doc.Document.Body.Items[idx].(type) {
		case *docx.Paragraph:
			for _, paragraphChildren := range element.Children {
				switch paragraphElement := paragraphChildren.(type) {
				case *docx.Run:
					if addSpacesBeetweenImageText && EF.checkHaveDrawing(paragraphElement) && EF.IsHaveEmptySpace(idx+2, &EF.EvertDoc.Doc.Document.Body.Items) {
						EF.AddSpace(idx+1, &EF.EvertDoc.Doc.Document.Body.Items)
						itemsLen++
					}

					if formatImagDescription && EF.checkHaveDrawing(paragraphElement) {
						EF.CapitalizePictureSentence(&EF.EvertDoc.Doc.Document.Body.Items[idx+1])
					}
				}
			}
		}
	}
}

func (EF *EvertFormatter) AddSpace(idx int, elements *[]interface{}) {
	buff := make([]interface{}, len((EF.EvertDoc.Doc.Document.Body.Items)[idx+1:]))
	copy(buff, (EF.EvertDoc.Doc.Document.Body.Items)[idx+1:])
	EF.EvertDoc.Doc.Document.Body.Items = (EF.EvertDoc.Doc.Document.Body.Items)[:idx+1]
	emptyParagraph := &docx.Paragraph{}
	EF.EvertDoc.Doc.Document.Body.Items = append(EF.EvertDoc.Doc.Document.Body.Items, emptyParagraph)
	EF.EvertDoc.Doc.Document.Body.Items = append(EF.EvertDoc.Doc.Document.Body.Items, buff...)
}

func (EF *EvertFormatter) IsHaveEmptySpace(idx int, elements *[]interface{}) bool {
	for _, el := range (EF.EvertDoc.Doc.Document.Body.Items)[idx].(*docx.Paragraph).Children {
		switch run := el.(type) {
		case *docx.Run:
			if run.Children == nil {
				return false
			}
		}
	}
	return true
}

func (EF *EvertFormatter) CapitalizePictureSentence(element *interface{}) {
	switch pr := (*element).(type) {
	case *docx.Paragraph:
		for _, parg := range pr.Children {
			switch run := parg.(type) {
			case *docx.Run:
				for _, text := range run.Children {
					switch text := text.(type) {
					case *docx.Text:
						fmt.Println(text.Text)
						r, size := utf8.DecodeRuneInString(text.Text)
						text.Text = string(unicode.ToUpper(r)) + text.Text[size:]
					}
				}
			}
		}
	}
}

func (EF *EvertFormatter) checkHaveDrawing(elements *docx.Run) bool {
	for _, runChildren := range elements.Children {
		switch runChildren.(type) {
		case *docx.Drawing:
			return true
		}
	}

	return false
}

func NewFormatter(doc *evert.EvertDoc) *EvertFormatter {
	return &EvertFormatter{
		EvertDoc: doc,
	}
}
