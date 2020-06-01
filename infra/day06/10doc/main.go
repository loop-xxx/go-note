package main

import (
	"fmt"

	"github.com/unidoc/unioffice/document"
)

func main() {
	if doc, err := document.Open("tempest.docx"); err == nil {
		paragraphs := []document.Paragraph{}
		for _, p := range doc.Paragraphs() {
			paragraphs = append(paragraphs, p)
		}

		// This sample document uses structured document tags, which are not common
		// except for in document templates.  Normally you can just iterate over the
		// document's paragraphs.
		for _, sdt := range doc.StructuredDocumentTags() {
			for _, p := range sdt.Paragraphs() {
				paragraphs = append(paragraphs, p)
			}
		}

		for _, p := range paragraphs {
			for _, r := range p.Runs() {
				fmt.Println(r.Text())
				/**
				switch r.Text() {
				case "崔坤坤":
					r.ClearContent()
					r.AddText("loop")
					//r.AddBreak() 换行
				}
				*/
			}
		}
		doc.SaveToFile("edit-document.docx")
	} else {
		fmt.Println(err)
	}

}
