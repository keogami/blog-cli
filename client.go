package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

func APIPath(components ...string) (string, error) {
  baseURL, err := url.Parse(os.Getenv("BLOG_API_URL"))
  if err != nil {
    return "", err
  }

  p := path.Join(components...)

  return baseURL.ResolveReference(&url.URL{ Path: p }).String(), nil 
}

func PostJson(path string, data interface{}) (*http.Response, error) {
  r, w := io.Pipe()
  go func() {
    json.NewEncoder(w).Encode(data)
    w.Close()
  }()

  return http.Post(path, "application/json", r)
}
