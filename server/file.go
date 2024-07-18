package server

type FileInfo struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Sha1sum string `json:"sha1sum"`
}
