package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path"
)

var BLOG_API_URL = os.Getenv("BLOG_API_URL")

func APIPath(components ...string) string {
  return path.Join(append([]string{BLOG_API_URL}, components...)...)
}

func PostJson(path string, data interface{}) (*http.Response, error) {
  r, w := io.Pipe()
  go json.NewEncoder(w).Encode(data)

  return http.Post(path, "application/json", r)
}
