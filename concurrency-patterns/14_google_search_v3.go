package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result
type SearchGen func(kind string) Search

func GetSearchSource(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func GetMultipleSearchSources(kind string, replica int) (replicas []Search) {
	for i := 0; i < replica; i++ {
		replicas = append(replicas, GetSearchSource(kind))
	}

	return
}

func First(query string, searchs []Search) Result {
	c := make(chan Result)

	for i := range searchs {
		go func(i int) { c <- searchs[i](query) }(i)
	}

	return <-c
}

func Google(query string) (results []Result) {
	var (
		webSources   = GetMultipleSearchSources("web", 3)
		imageSources = GetMultipleSearchSources("image", 3)
		videoSources = GetMultipleSearchSources("video", 3)
	)

	c := make(chan Result)

	go func() { c <- First(query, webSources) }()
	go func() { c <- First(query, imageSources) }()
	go func() { c <- First(query, videoSources) }()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("search timed out")
			return
		}
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
