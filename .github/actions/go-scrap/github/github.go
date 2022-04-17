package github

import (
	"fmt"
	"os"
	"strings"
)

func GetInput(name string) string {
	name = fmt.Sprintf("INPUT_%s", strings.ToUpper(name))
	val := os.Getenv(name)
	return strings.TrimSpace(val)
}
