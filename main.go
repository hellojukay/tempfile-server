package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 3456, "port")
	if !flag.Parsed() {
		flag.Parse()
	}
}

func main() {
	// create http server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
