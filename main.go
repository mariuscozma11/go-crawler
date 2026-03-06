package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "BotCrawler/1.0")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode == 400 {
		return "", fmt.Errorf("status code 400")
	}
	if res.Header.Get("content-type") != "text/html" {
		return "", fmt.Errorf("content type header is not text/html")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	stringBody := string(body)

	return stringBody, nil

}

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Printf("no website provided\n")
		os.Exit(1)
	}
	if len(os.Args[1:]) > 1 {
		fmt.Printf("too many arguments provided\n")
		os.Exit(1)
	}
	baseURL := os.Args[1]

	fmt.Printf("starting crawl of: %s\n", baseURL)
	html, err := getHTML(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(html + "\n")

}
