// search 包包含程序使用的主框架和逻辑
package search

import (
	"fmt"
	"log"
	"sync"
)

// a map of registered matchers for searching.
var matchers = make(map[string]Matcher)

// Register is called to register a matcher
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}

// Run performs the search logic.
func Run(searchTerm string) {
	// Retrieve the list of feeds to search through.
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal("RetrieveFeeds: ", err)
	}

	// Create an unbuffered channel to receive match results to display.
	results := make(chan *Result)

	// Setup a wait group so we can process all the feeds.
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	// Launch a goroutine for each feed to find the result.
	for _, feed := range feeds {
		// Retrieve a matcher to search
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Launch a goroutine to perform the search.
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// launch a goroutine to monitor all the work is done.
	go func() {
		// Wait for everything to be processed.
		waitGroup.Wait()
		// Close the channel to signal to the Display
		// function that we can exit the program.
		close(results)
	}()

	// Start displaying the results as they are available and
	// return after the final result is displayed.
	Display(results)
}

// Display writes results to the console window as they are
// received by the individual goroutines.
func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
