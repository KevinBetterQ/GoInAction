package search

import (
	"encoding/json"
	"os"
)

const dataFile = "example/data/data.json"

// Feed contains information we need to process a feed
type Feed struct {
	Name string `json:"site"`
	URL  string `json:"link"`
	Type string `json:"type"`
}

// RetriveFeeds reads and unmarshals the feed data file.
func RetrieveFeeds() ([]*Feed, error) {
	// Open the file
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Decode the file into a slice of pointers to Feed values.
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
