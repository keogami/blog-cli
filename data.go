package main

import (
	"os"
)

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
