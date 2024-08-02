package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/hellojukay/tempfile-server/event"

	"github.com/hellojukay/tempfile-server/config"
	"github.com/hellojukay/tempfile-server/server"
)

var (
	version bool
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.IntVar(&config.Port, "port", 3456, "port")
	flag.StringVar(&config.Dir, "dir", "./", "server directory")
	flag.BoolVar(&version, "version", false, "show version")
	flag.DurationVar(&event.Duration, "expiretime", 3600*24, "expire time")
	flag.StringVar(&config.ExternalURL, "external-url", "http://127.0.0.1:3456", "server external url")
	if !flag.Parsed() {
		flag.Parse()
	}
	if version {
		info, ok := debug.ReadBuildInfo()
		if ok {
			println(info.Main.Version, info.Main.Sum)
		}
		os.Exit(0)
	}
}

func main() {
	// create http server
	http.Handle("/", server.NewFileServer(config.Dir))
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
