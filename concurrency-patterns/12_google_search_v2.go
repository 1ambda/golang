package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

func getSearchSource(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result) {
	var (
		Web   = getSearchSource("web")
		Image = getSearchSource("image")
		Video = getSearchSource("video")
	)

	c := make(chan Result)

	go func() { c <- Web(query) }()
	go func() { c <- Video(query) }()
	go func() { c <- Image(query) }()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}

	return
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// do search
	start := time.Now()
	results := Google("golang")
	elasped := time.Since(start)

	fmt.Println(results)
	fmt.Println(elasped)
}
