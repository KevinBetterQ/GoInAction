package matcher

import "GoInAction/sample1/search"

type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	search.Register("default", matcher)
}

func (m defaultMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	return nil, nil
}
