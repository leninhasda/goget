package parser

import (
	"fmt"
	"log"
	neturl "net/url"
	"strings"
)

// ArgsURL parses and escape urls from command line args
func ArgsURL(args []string) []string {
	urls := args

	firstArg := args[0]
	if firstArg == "-parallel" { // ignore first 2 if it's flag
		urls = args[2:]
	}

	parsedURLs := []string{}
	for _, u := range urls {
		parsed, err := parseURL(u)
		if err != nil {
			log.Printf("warn: invalid url %s", u) // ignore invalid url
			continue
		}
		parsedURLs = append(parsedURLs, parsed)
	}

	return parsedURLs
}

func parseURL(url string) (string, error) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url // add schema if its missing
	}

	_, err := neturl.ParseRequestURI(url)
	if err != nil || !strings.Contains(url, ".") {
		return "", fmt.Errorf("invalid url %s", url) // ignore invalid url
	}

	return url, nil
}
