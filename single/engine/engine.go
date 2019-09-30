package engine

import (
	"log"

	"github.com/Felyne/crawler/single/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s\n", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("fetch url: %s error: %v\n", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
