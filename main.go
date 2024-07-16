package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/hellojukay/tempfile-server/server"
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
	http.Handle("/", server.NewFileServer("/tmp/"))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
