package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/hellojukay/tempfile-server/event"

	"github.com/hellojukay/tempfile-server/config"
	"github.com/hellojukay/tempfile-server/server"
)

var (
	version bool
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.IntVar(&config.Port, "port", 3456, "Port")
	flag.StringVar(&config.Dir, "dir", "./", "Server directory")
	flag.BoolVar(&version, "version", false, "Show version")
	flag.DurationVar(&event.Duration, "expiretime", time.Duration(time.Hour*8), "Expire time")
	flag.StringVar(&config.ExternalURL, "external-url", "http://127.0.0.1:3456", "Server external url")
	flag.StringVar(&config.AccessKeyID, "access-key", "", "S3 server access key")
	flag.StringVar(&config.AccessKeySecret, "secret-key", "", "S3 server secret key")
	flag.StringVar(&config.BucketName, "bucket", "upload", "S3 server bucket name")
	flag.StringVar(&config.S3EndPoint, "endpoint", "127.0.0.1:9000", "S3 server endpoint")
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
	//http.Handle("/", server.NewFileServer(config.Dir))
	http.Handle("/", &server.MinIOServer{})

	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
