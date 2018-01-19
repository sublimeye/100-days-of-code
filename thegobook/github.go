package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url`
}

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed. Query '%v': Status Code %v", q, resp.StatusCode)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}

// Copied from stackoverflow (too much shit in here, why Golang doesn't have this natively?)
func TimeDiff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

const temp1 = `{{.TotalCount}} issues:
{{range .Items}}----------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func runSearchIssues() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// sort by CreatedAt
	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].CreatedAt.Before(result.Items[j].CreatedAt)
	})

	// lineOutput(result)
	// templateOutput(result)
	htmlTemplateOutput(result)
}

func lineOutput(result *IssuesSearchResult) {
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		year, month, _, _, _, _ := TimeDiff(item.CreatedAt, time.Now())
		if year > 1 {
			fmt.Printf("[More than year] #%-5d %v %9.9s %.55s\n", item.Number, item.CreatedAt, item.User.Login, item.Title)
		} else if year <= 1 && month >= 1 {
			fmt.Printf("[Less than year] #%-5d %v %9.9s %.55s\n", item.Number, item.CreatedAt, item.User.Login, item.Title)
		} else {
			fmt.Printf("[Less than month] #%-5d %v %9.9s %.55s\n", item.Number, item.CreatedAt, item.User.Login, item.Title)
		}
	}
}

func templateOutput(result *IssuesSearchResult) {
	report := template.Must(template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(temp1))

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

const htmlTemplate = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

func htmlTemplateOutput(result *IssuesSearchResult) {
	var issueList = template.Must(template.New("issuelisthtml").Parse(htmlTemplate))
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func main() {
	runSearchIssues()
}
