// search 包包含程序使用的主框架和逻辑
package search

import (
	"fmt"
	"log"
	"sync"
)

// 创建一个匹配器映射，用于将各个匹配器进行注册
var matchers = make(map[string]Matcher)

// Register 用于注册匹配器
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}

// Run 执行搜索逻辑
func Run(searchTerm string) {
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal("RetrieveFeeds: ", err)
	}

	// 创建一个无缓冲的通道，接收匹配后的结果
	results := make(chan *Result)

	// 构造一个 waitGroup 并设置，以便处理所有的数据源
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个 goroutine 来查找结果
	for _, feed := range feeds {
		// 获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 启动一个 goroutine 来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个 goroutine 来监控是否所有工作都做完了
	go func() {
		// 等候所有任务完成
		waitGroup.Wait()
		// 以关闭 results 的方式来通知 Display 函数退出
		close(results)
	}()

	// 显示返回的结果，并在最后一个结果显示完成后返回
	Display(results)
}

// Display 从每个单独的 goroutine 中接收结果后进行显示
func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
