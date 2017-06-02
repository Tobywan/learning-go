package ch1

import (
	"fmt"
	"log"
	"net/http"
)

func doServer() {
	http.HandleFunc("/", simplehandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func simplehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
