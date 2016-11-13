package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	http.HandleFunc("/search", handleSearch)
	fmt.Println("serving on http://localhost:7777/search")
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

type Result struct {
	Title, URL string
}

type templateData struct {
	Results []Result
	Elasped time.Duration
}

// handleSearch handles URLs like "/search?q=golang" by running a
// Google search for "golang" and writing the results as HTML to w.
func handleSearch(w http.ResponseWriter, req *http.Request) {
	log.Println("serving", req.URL)

	query := req.FormValue("q")

	if query == "" {
		http.Error(w, `missing "q" URL parameter`, http.StatusBadRequest)
		return
	}

	start := time.Now()
	results, err := Search(query)
	elasped := time.Since(start)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var resultsTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
  <ol>
  {{range .Results}}
    <li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
  {{end}}
  </ol>
  <p>{{len .Results}} results </p>
</body>
</html>
`))

	if err := resultsTemplate.Execute(w, templateData{
		Results: results,
		Elasped: elasped,
	}); err != nil {
		log.Print(err)
		return
	}
}

func Search(query string) ([]Result, error) {
	u, err := url.Parse("https://ajax.googleapis.com/ajax/services/search/web?v=1.0")
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("q", query)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var jsonResponse struct {
		ResponseData struct {
			Results []struct {
				TitleNoFormatting, URL string
			}
		}
	}
	if err := json.NewDecoder(resp.Body).Decode(&jsonResponse); err != nil {
		return nil, err
	}

	// Extract the Results from jsonResponse and return them.
	var results []Result
	for _, r := range jsonResponse.ResponseData.Results {
		results = append(results, Result{Title: r.TitleNoFormatting, URL: r.URL})
	}
	return results, nil
}
