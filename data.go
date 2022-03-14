package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type APIError struct {
  Message string `json:"message"`
}

func ParseError(r io.Reader) (APIError, error) {
  var e APIError
  err := json.NewDecoder(r).Decode(&e)
  if err != nil {
    return e, err
  }

  return e, nil
}

func (em APIError) Error() string {
  return fmt.Sprintf("api error: %s", em.Message)
}

type BlogPost struct {
  Meta BlogMeta `json:"meta"`
  Content string `json:"content"`
}

type BlogMeta struct {
  Slug string `json:"slug"`
  Title string `json:"title"`
  Group interface{} `json:"group"`
  PostTime time.Time `json:"postTime"`
  Summary string `json:"summary"`
  Tags []string `json:"tags"`
}

func LoadBlogPost(path string) (BlogPost, error) {
  f, err := os.Open(path)
  if err != nil {
    return BlogPost{}, err
  }
  defer f.Close()

  var p BlogPost
  err = json.NewDecoder(f).Decode(&p)
  if err != nil {
    return BlogPost{}, err
  }

  return p, nil
}

func (p BlogPost) SaveBlogPost(path string) error {
  f, err := os.Create(path)
  if err != nil {
    return err
  }

  return json.NewEncoder(f).Encode(p)
}

type CreateParams struct {
  Meta
  Content string `json:"content" yaml:"content"`
}

func LoadMeta() (Meta, error) {
  f, err := os.Open("meta.yml")
  if err != nil {
    return Meta{}, err
  }
  defer f.Close()

  return MetaFromYaml(f)
}

func LoadContent() (string, error) {
  content, err := os.ReadFile("content.md")
  return string(content), err
}

func NewCreateParams() (CreateParams, error) {
  m, err := LoadMeta()
  if err != nil {
    return CreateParams{}, err
  }

  c, err := LoadContent()
  if err != nil {
    return CreateParams{}, err
  }

  return CreateParams{
    Meta {
      Title: m.Title,
      Summary: m.Summary,
      Tags: m.Tags,
    },
    c,
  }, nil
}
