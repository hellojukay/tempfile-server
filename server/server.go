package server

import (
	"net/http"
	"os"
	"path/filepath"
)

func NewFileServer(dir string) http.Handler {
	return FileServer(Dir(abs(dir)), abs(dir))
}

func abs(dir string) string {
	if filepath.IsAbs(dir) {
		return dir
	}
	pwd, _ := os.Getwd()
	return filepath.Join(pwd, dir)
}
