package engine

import (
	"github.com/Kirk-Wang/Hello-Gopher/history/17.6/crawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil
}
