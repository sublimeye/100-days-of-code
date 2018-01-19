// started 8:50
package main

import "time"
import "text/template"

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

report, err := template.New("report").
Funcs(template.FuncMap{"daysAgo": daysAgo}).
Parse(temp1)
if err != nil {
  log.Fatal(err)
}
