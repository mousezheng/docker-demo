package main

import (
	"fmt"
	"github.com/expectedsh/go-sonic/sonic"
)

func main() {

	ingester, err := sonic.NewIngester("127.0.0.1", 14911, "SecretPassword")
	if err != nil {
		panic(err)
	}

	// I will ignore all errors for demonstration purposes

	_ = ingester.BulkPush("movies", "general", 3, []sonic.IngestBulkRecord{
		{"id:1", "今天是晴天"},
		{"id:2", "今天是雨天"},
		{"id:3", "今天是阴天"},
	}, sonic.LangAutoDetect)
	if err != nil {
		panic(err)
	}

	search, err := sonic.NewSearch("localhost", 14911, "SecretPassword")
	if err != nil {
		panic(err)
	}

	results, _ := search.Query("movies", "general", "晴天", 10, 0, sonic.LangAutoDetect)

	fmt.Println(results)
}