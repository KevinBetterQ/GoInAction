package search

import (
	"encoding/json"
	"os"
)

const dataFile = "sample1/data/data.json"

// Feed 描述每个数据源的信息
type Feed struct {
	Name string `json:"site"`
	URL  string `json:"link"`
	Type string `json:"type"`
}

// RetriveFeeds 获取需要搜索的数据源列表
func RetrieveFeeds() ([]*Feed, error) {
	// 打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// 读取文件解码到一个切片
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	// 返回结果
	return feeds, err
}
