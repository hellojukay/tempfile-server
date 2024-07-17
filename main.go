package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/hellojukay/tempfile-server/server"
)

var (
	port int
	dir  string
)

func init() {
	flag.IntVar(&port, "port", 3456, "port")
	flag.StringVar(&dir, "dir", "./", "server directory")
	if !flag.Parsed() {
		flag.Parse()
	}
}

func main() {
	// create http server
	http.Handle("/", server.NewFileServer(dir))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
