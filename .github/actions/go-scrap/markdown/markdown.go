package markdown

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/n4ze3m/ycombinator-jobs/github"
)

const (
	README_PATH = "./README.md"
)

func Readme(table string) {
	// open file and get data
	file, err := os.OpenFile(README_PATH, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// file size
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	data := make([]byte, stat.Size())
	file.Read(data)
	start := strings.Index(string(data), "<!-- JOBS:START -->")
	end := strings.Index(string(data), "<!-- JOBS:END -->")

	if start == -1 || end == -1 {
		return
	}


	file.Seek(0, 0)
	// remove all content between start and end
	file.WriteAt(data[:start], 0)
	file.WriteAt(data[end:], int64(start))
	// write new content
	newData := fmt.Sprintf("%s\n%s\n", "<!-- JOBS:START -->", table)
	file.WriteAt([]byte(newData), int64(start))
	file.WriteAt(data[end:], int64(start+len(newData)))


	fmt.Println("Updated README.md")
}


func Push() {
	repo := os.Getenv("GITHUB_REPOSITORY")
	githubToken := github.GetInput("gh_token")

	url := fmt.Sprintf("%s@github.com/%s.git", githubToken, repo)
	// set origin-url
	exec.Command("git", "remote", "set-url", "origin", url).Run()
	// set global user.name
	exec.Command("git", "config", "--global", "user.name", "n4ze3m").Run()
	// set global user.email
	exec.Command("git", "config", "--global", "user.email", "mock@n4ze3m.site").Run()
	// add readme
	exec.Command("git", "add", README_PATH).Run()
	// commit
	exec.Command("git", "commit", "-m", "Update README.md").Run()
	// push
	exec.Command("git", "push", "origin", "main").Run()

	fmt.Println("Pushed to github")
}