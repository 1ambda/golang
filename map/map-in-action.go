package main

import (
	_ "fmt"
	"sort"
	"sync"
)

type Key struct {
	Path, Country string
}

// Ref: https://blog.golang.org/go-maps-in-action
func main() {
	/**
	 * Map types are reference types, like pointers or slices, and
	 * so the value of `m` below is `nil` It doens't point to an initialized map.
	 * Attemping to write to a nil map will cause runtime panic.
	 *
	 * The `make` function allocates and initializes a hash map and
	 * returns map value that points to it.
	 */

	var m map[string]string
	m = make(map[string]string)
	m["route"] = "1"
	test(m["route"] == "1")

	n := len(m)
	test(n == 1)

	// The `delete` function doesn't return anything, and will do othing
	// If the specified key doesn't exist
	delete(m, "route")
	_, ok := m["route"]
	test(ok == false)

	m["root"] = "admin"

	for key, value := range m {
		test(key == "root")
		test(value == "admin")
	}

	// Initializing a map with data
	commits := map[string]int{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	test(len(commits) == 4)

	empty := map[string]int{}
	test(len(empty) == 0)

	// Language spec defines the possible keys for map, but in short, comparable types are
	// boolean, numeric, string, pointer, channel, interface types, and structs or arrays
	// Not slice, maps, functions! these types can not be compared using `==`
	// But structs can be used as the key

	hits := make(map[string]map[string]int)
	test(hits["/doc/"]["au"] == 0)

	// We need to create an map for unvisited url
	add := func(m map[string]map[string]int, path, country string) {
		mm, ok := m[path]
		if !ok {
			mm = make(map[string]int)
			m[path] = mm
		}

		mm[country]++
	}

	add(hits, "/doc/", "au")
	test(hits["/doc/"]["au"] == 1)

	// On the other hand, a design that uses a single map with a struct key can help us!

	hits2 := make(map[Key]int)
	hits2[Key{"/", "vn"}]++
	test(hits2[Key{"/ref/spec", "ch"}] == 0)

	// Important:
	// Maps are not safe for concurrent use: https://golang.org/doc/faq#atomic_maps

	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	counter.RLock()
	count := counter.m["some_key"]
	counter.RUnlock()
	test(count == 0)

	// Important:
	// The iteration order is not specified and is not guaranteed to be the same

	var unordered map[int]string
	var keys []int
	for k := range unordered {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	test(len(unordered) == 0)
}

func test(expr bool) {
	if !expr {
		panic(0)
	}
}

func use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}
