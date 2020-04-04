package main

import (
	"log"
	"os"

	_ "GoInAction/sample1/matcher"
	"GoInAction/sample1/search"
)

// init is called prior to main.
func init() {
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// 使用特定的项做搜索
	// Perform the search for the specified term.
	search.Run("president")
}
