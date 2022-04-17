package main

import (
	"fmt"
	"os"
	"strings"
)

func getInput(name string) string {
	name = fmt.Sprintf("INPUT_%s", strings.ToUpper(name))
	val := os.Getenv(name)
	return strings.TrimSpace(val)
}


func main() {
	fmt.Println(getInput("committer_username"))
	fmt.Println(getInput("committer_email"))
	fmt.Println(getInput("commit_message"))
}
