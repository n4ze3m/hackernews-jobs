package fetch

import (
	"fmt"
	"regexp"
	"strings"
	"time"

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
			dateTime, isExist := el.DOM.Next().Find(".age").Attr("title")
			t := ""
			if !isExist {
				t = strings.TrimSpace(el.DOM.Next().Text())
			} else {
				tm, err:= time.Parse("2006-01-02T15:04:05", dateTime)
				if err != nil {
					fmt.Println(err)
				}
				t = tm.Format("1/2/2006 03:04 PM")
			}

			if !regexp.MustCompile(`^(http|https)://`).MatchString(link) {
				link = "https://news.ycombinator.com/" + link
			}

			jobs = append(jobs, Jobs{
				Title: title,
				Link:  link,
				Time:  t,
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
