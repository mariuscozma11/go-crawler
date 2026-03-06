package main

import (
	"net/url"
)

func normalizeURL(inputURL string) (string, error) {
	URL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}
	formattedURL := URL.Host + URL.Path
	return formattedURL, nil
}
