package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	TARGET_URL = "https://news.ycombinator.com/jobs"
)

func Run() {
	client := &http.Client{}
	
	req, err := http.NewRequest("GET", TARGET_URL, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
