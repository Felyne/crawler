package engine

import (
	"log"

	"github.com/Felyne/crawler/concurrent/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch url: %s error: %v\n", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
