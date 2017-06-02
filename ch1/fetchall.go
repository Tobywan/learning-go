package ch1

import (
	"fmt"
	"io"
	//	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func doFetch() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch, start.Unix()) // a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, timeStamp int64) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	f, err := getOutPutFile(url, start.Unix())
	defer f.Close()
	if err != nil {
		ch <- fmt.Sprintf("while creating file: %v", err)
		return
	}

	//	nbytes, err := f.Write(resp.Body)
	nbytes, err := io.Copy(f, resp.Body)

	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func getOutPutFile(url string, timestamp int64) (*os.File, error) {

	parts := []string{"/home/toby/junk/",
		strconv.FormatInt(timestamp, 10),
		"/"}

	dirname := strings.Join(parts, "")
	os.MkdirAll(dirname, 0733)
	filename := strings.Join([]string{dirname, strings.Split(url, "//")[1]}, "")

	return os.Create(filename)

}
