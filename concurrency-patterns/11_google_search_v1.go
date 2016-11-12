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

	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))

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
