package main

import (
	"log"
	"os"

	_ "GoInAction/sample1/matcher"
	"GoInAction/sample1/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	// 使用特定的项做搜索
	search.Run("test")
}
