package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/urfave/cli/v2"
)

func Post(c *cli.Context) error {
  params, err := NewCreateParams()
  if err != nil {
    return cli.Exit(fmt.Errorf("Couldn't load post parameters: %w", err), 2)
  }

  p, err := APIPath("/blog")
  if err != nil {
    return cli.Exit(fmt.Errorf("Couldn't construct API route: %w", err), 4)
  }

  res, err := PostJson(p, params)
  if err != nil {
    return cli.Exit(fmt.Errorf("Couldn't send post parameters: %w", err), 3)
  }
  defer res.Body.Close()

  status := res.StatusCode
  b := new(strings.Builder)

  io.Copy(b, res.Body)

  fmt.Println(status)
  fmt.Println(b.String())
  return nil
}
