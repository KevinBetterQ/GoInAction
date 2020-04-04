package search

import "log"

// Result 描述搜索的结果
type Result struct {
	Field   string
	Content string
}

// Matcher 定义了要实现的新搜索类型的行为
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 搜索数据源的数据，并将匹配结果输出到 results 通道
func Match(matcher Matcher, feed *Feed, term string, results chan *Result) {
	// 对特定的搜索器执行搜索
	searchResults, err := matcher.Search(feed, term)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}
