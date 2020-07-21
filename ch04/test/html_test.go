package test

import (
	"html/template"
	"log"
	"testing"
	"time"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

func TestHtml(t *testing.T) {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", report)
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
