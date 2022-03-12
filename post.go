package main

import (
	"encoding/json"
	"fmt"

	"github.com/urfave/cli/v2"
)

func Post(c *cli.Context) error {
  p, err := NewCreateParams()
  if err != nil {
    return cli.Exit(fmt.Errorf("Couldn't load post parameters: %w", err), 2)
  }

  d, _ := json.MarshalIndent(p, "", "  ")
  fmt.Println(string(d))
  return nil
}
