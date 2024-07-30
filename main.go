package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/hellojukay/tempfile-server/config"
	"github.com/hellojukay/tempfile-server/server"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.IntVar(&config.Port, "port", 3456, "port")
	flag.StringVar(&config.Dir, "dir", "./", "server directory")
	flag.StringVar(&config.ExternalURL, "external-url", "http://127.0.0.1:3456", "server external url")
	if !flag.Parsed() {
		flag.Parse()
	}
}

func main() {
	// create http server
	http.Handle("/", server.NewFileServer(config.Dir))
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
