# tempfile-server

A tempfile server, designed for CI system, save file like logs , archive files and and artifacts, files will be shared between CI or mutil deparments。

Working in progress, not ready to use.

# Install
```bash
go install github.com/yubo/tempfile-server@latest
``` 

# Usage
Command line usage help
```bash
$ ./tempfile-server -h
Usage of ./tempfile-server:
  -access-key string
        S3 server access key
  -bucket string
        S3 server bucket name (default "upload")
  -dir string
        Server directory (default "./")
  -expiretime duration
        Expire time (default 8h0m0s)
  -external-url string
        Server external url (default "http://127.0.0.1:3456")
  -port int
        Port (default 3456)
  -secret-key string
        S3 server secret key
  -version
        Show version
```
How to upload file to server.
```bash
$ curl -T tempfile-server http://localhost:3456/aaa/bbb/c/
{
    "name": "tempfile-server",
    "path": "/aaa/bbb/c/tempfile-server",
    "sha1sum": "d61827094d019e123aefe57588916e85c79ce99f",
    "download_url": "http://127.0.0.1:3456/aaa/bbb/c/tempfile-server"
}
```

TODO : 
- [x] support upload file to disk.
- [ ] support s3.
- [ ] permession control.
- [ ] buket control.


