package engine

import (
	"log"

	"github.com/Felyne/crawler/distributed/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s\n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch url: %s error: %v\n", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil
}
