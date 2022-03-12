package main

import (
	"fmt"
	"os"
  _ "embed"

	"github.com/urfave/cli/v2"
)

//go:embed init_data/meta.yml
var FileMetaYML []byte

//go:embed init_data/content.md
var FileContentMD []byte

var files = map[string][]byte {
  "meta.yml": FileMetaYML,
  "content.md": FileContentMD,
}

func Init(c *cli.Context) error {
  for path, data := range files {
    f, err := os.Create(path)
    if err != nil {
      return cli.Exit(fmt.Errorf("Couldn't create [%s]: %w", path, err), int(ErrFile))
    }

    _, err = f.Write(data)
    if err != nil {
      return cli.Exit(fmt.Errorf("Couldn't create [%s]: %w", path, err), int(ErrFile))
    }
  }
  fmt.Println("New Blog Post entry initialized.")
  return nil
}
