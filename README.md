# tempfile-server

A tempfile server, designed for CI system, save file like logs , archive files and and artifacts, files will be shared between CI or mutil deparments。

Working in progress, not ready to use.

# Usage
How to upload file to server.
```bash
# directory will auto created if not exists.
curl -T file http://localhost:3456/..{directory}/.../{filename}
```

TODO : A tmepfile server , upload file to server and return download url, and files will be auto deleted after some days。


