package search

import "log"

// Result contains the result of a search
type Result struct {
	Field   string
	Content string
}

// Matcher defined the behavior required by types that want
// to implement a new search type
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is launched as a goroutine for each individual feed
// to run searches concurrently.
func Match(matcher Matcher, feed *Feed, term string, results chan *Result) {
	// Perform the search against the specified matcher.
	searchResults, err := matcher.Search(feed, term)
	if err != nil {
		log.Println(err)
		return
	}

	// Write the results to the channel.
	for _, result := range searchResults {
		results <- result
	}
}
