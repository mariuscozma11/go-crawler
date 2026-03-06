package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getFirstParagraphFromHTML(html string) string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)

	}
	firstParagraph := doc.Find("main").Find("p").First().Text()
	if firstParagraph == "" {

		firstParagraph = doc.Find("body").Find("p").First().Text()
	}
	return firstParagraph
}
