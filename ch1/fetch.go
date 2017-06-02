package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	httpPrefix = "http://"
)

func main_fetch() {

	for _, url := range os.Args[1:] {
		fmt.Println("\n=====================")
		readUrl(url)
	}

}

func readUrl(url string) {

	prefix := ""
	if !strings.HasPrefix(url, httpPrefix) {
		prefix = httpPrefix
	}

	resp, err := http.Get(prefix + url)
	handleErr(err, "fetching")

	fmt.Printf("Status:%d:%s\n", resp.StatusCode, resp.Status)

	_, err = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	handleErr(err, "copying")

}

func handleErr(err error, action string) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "%s: %v\n", action, err)
	os.Exit(1)

}
