package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)

	}
	heading := doc.Find("h1").Text()
	if heading == "" {
		heading = doc.Find("h2").Text()
	}

	return heading
}
