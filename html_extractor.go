package main

import (
	"fmt"
	"log"
	"net/url"
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
func getFirstParagraphFromHTML(html string) string {

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)

	}
	firstParagraph := doc.Find("main").Find("p").First().Text()
	return firstParagraph
}
func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %w", err)
	}

	var urls []string
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			return
		}
		href = strings.TrimSpace(href)
		if href == "" {
			return
		}

		u, err := url.Parse(href)
		if err != nil {
			fmt.Printf("couldn't parse href %q: %v\n", href, err)
			return
		}

		resolved := baseURL.ResolveReference(u)
		urls = append(urls, resolved.String())
	})

	return urls, nil
}
func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %w", err)
	}

	var imageURLs []string
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if !ok || strings.TrimSpace(src) == "" {
			return
		}

		u, err := url.Parse(src)
		if err != nil {
			fmt.Printf("couldn't parse src %q: %v\n", src, err)
			return
		}

		absolute := baseURL.ResolveReference(u)
		imageURLs = append(imageURLs, absolute.String())
	})

	return imageURLs, nil
}

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	return PageData{}
}
