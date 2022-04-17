package fetch

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/n4ze3m/ycombinator-jobs/markdown"
)

const (
	TARGET_URL = "https://news.ycombinator.com/jobs"
	HEADER     = `| id 	| Title 	| Posted On 	|
|---	|---	|---	|`
)

type Jobs struct {
	Title string
	Link  string
	Time  string
}

func Run() {
	c := colly.NewCollector()
	markDownTable := []string{}
	jobs := []Jobs{}

	c.OnHTML(".itemlist", func(e *colly.HTMLElement) {
		e.ForEach(".athing", func(_ int, el *colly.HTMLElement) {
			title := el.ChildText(".title")
			link := el.ChildAttr(".title a", "href")
			time := strings.TrimSpace(el.DOM.Next().Text())

			if !regexp.MustCompile(`^(http|https)://`).MatchString(link) {
				link = "https://news.ycombinator.com/" + link
			}

			jobs = append(jobs, Jobs{
				Title: title,
				Link:  link,
				Time:  time,
			})
		})
	})

	c.Visit(TARGET_URL)

	markDownTable = append(markDownTable, HEADER)

	for i, job := range jobs {
		markDownTable = append(markDownTable, fmt.Sprintf("| %d | [%s](%s) | %s |", i+1, job.Title, job.Link, job.Time))
	}


	markdown.Readme(strings.Join(markDownTable, "\n"))
	markdown.Push()
}
