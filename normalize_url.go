package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	//parse the URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// Convert to lowercase
	host := strings.ToLower(parsedURL.Host)

	//remove port number
	if i := strings.IndexByte(host, ':'); i != -1 {
		host = host[:i]
	}

	//remove www. prefix if present
	if strings.HasPrefix(host, "www.") {
		host = host[4:]
	}

	//get path and remove trailing slash
	path := parsedURL.Path
	if path != "" && strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}

	normalized := host + path

	return normalized, nil
}
