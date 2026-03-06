package main

import (
	"log"
	"net/url"
)

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	var pageData PageData
	url, err := url.Parse(pageURL)
	if err != nil {
		log.Fatal(err)
	}

	pageData.URL = pageURL
	pageData.Heading = getHeadingFromHTML(html)
	pageData.FirstParagraph = getFirstParagraphFromHTML(html)
	pageData.OutgoingLinks, err = getURLsFromHTML(html, url)
	if err != nil {
		log.Fatal(err)
	}
	pageData.ImageURLs, err = getImagesFromHTML(html, url)
	if err != nil {
		log.Fatal(err)
	}
	return pageData
}

//   {URL:https://crawler-test.com Heading:Test Title FirstParagraph:This is the first paragraph. OutgoingLinks:[https://crawler-test.com/link1] ImageUR
//Ls:[https://crawler-test.com/image1.jpg]}
//	 {URL:https://crawler-test.com Heading:Test Title FirstParagraph: OutgoingLinks:[https://crawler-test.com/link1] ImageURLs:[https://crawler-test
//	.com/image1.jpg]}
