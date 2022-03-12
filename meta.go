package main

import (
	"io"

	"gopkg.in/yaml.v3"
)

type Meta struct {
  Title string `json:"title" yaml:"title"`
  Summary string `json:"summary" yaml:"summary"`
  Tags []string `json:"tags" yaml:"tags"`
}

func (m *Meta) ToYaml(w io.Writer) (error) {
  return yaml.NewEncoder(w).Encode(m)
}

func MetaFromYaml(w io.Reader) (m Meta, err error) {
  err = yaml.NewDecoder(w).Decode(&m)
  return
}
