package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
	"strings"
)

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// enhanced version
func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	w.Header().Set("Content-Type", "text/plain")
	if strings.HasSuffix(path, "@golang.org") {
		name := strings.TrimSuffix(path, "@golang.org")
		fmt.Fprintf(w, "Hello gopher %s\n", name)
		return
	}

	fmt.Fprintf(w, "Hello dear %s\n", path)
}

var re = regexp.MustCompile("^(.+)@golang.org$")

func handler2(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	match := re.FindAllStringSubmatch(path, -1)

	if match != nil {
		fmt.Fprintf(w, "Hello gopher %s\n", match[0][1])
		return
	}

	fmt.Fprintf(w, "Hello dear %s\n", path)
}
