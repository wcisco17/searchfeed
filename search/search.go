package search

import (
	"log"
	"sync"
)

// A map of registered matchers for searching
var matchers = make(map[string]Matcher)

// Run performs the search logic
func Run(searchTerm string, site string) {
	// Retrieve the list of feeds to search through.
	feeds, err := RetrieveFeeds()

	if err != nil {
		log.Fatal(err)
	}

	// Create an unbuffered channel to receive match results
	results := make(chan *Result)

	// Set up a wait group so we can process all the feeds
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait a while.
	waitGroup.Add(len(feeds))
	// Launch a goroutines for each feed to find the results.
	for _, feed := range feeds {
		// Retrieve a matcher for the search
		matcher, exists := matchers[feed.Type]

		if !exists {
			matcher = matchers["default"]
		}

		// Launch the goroutine to perform the search.
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results, site)
			waitGroup.Done()
		}(matcher, feed)
	}

	// Launch a goroutine to monitor when all the work is done.
	go func() {
		// Wait for everything to complete.
		waitGroup.Wait()

		// Close the channel to signal to the Display
		// function that we can exit the program.
		close(results)
	}()

	// Start displaying results as they are available.

	// Return after the final result is displayed.
	Display(results)
}

// Register is called to register a matcher for use by the program.
// This will run first since it's called in the init function of default.go
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
