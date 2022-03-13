package api

import (
	"microstorage"
	"net/http"
	"path/filepath"
)

func init() {
	go func() {
		http.HandleFunc("/api/files/list", handleFileList)
	}()
}

func handleFileList(w http.ResponseWriter, r *http.Request) {
	globs, _ := filepath.Glob(microstorage.Config.App.StoragePath + "/data/*/*")
	out := ""
	w.Header().Set("Content-type", "text/html")
	for _, fname := range globs {
		bn := filepath.Base(fname)
		out = out + "<div><a href='/get/" + bn + "'>" + bn + "</a></div>"
	}
	w.Write([]byte(out))
}
