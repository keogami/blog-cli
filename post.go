package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/urfave/cli/v2"
)

func Post(c *cli.Context) error {
  bp, err := LoadBlogPost(".blogpost.json")
  if err == nil {
    return cli.Exit(fmt.Errorf("Entry was already posted at /%s. Did you mean `put`?", bp.Meta.Slug), 9)
  }


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

  if res.StatusCode != http.StatusCreated {
    return HandleErrorStatus(res, map[int]string{
      http.StatusBadRequest: "There's something wrong with this client",
      http.StatusInternalServerError: "Server Error: %w",
    })
  }

  err = json.NewDecoder(res.Body).Decode(&bp)
  if err != nil {
    return cli.Exit(fmt.Errorf("Couldn't parse the Blog Post returned by the server: %w", err), 8)
  }

  err = bp.SaveBlogPost(".blogpost.json")
  if err != nil {
    return cli.Exit(fmt.Errorf("Couldn't save the response sent by the server: %w", err), 9)
  }

  fmt.Println("Entry was posted:", "/" + bp.Meta.Slug)
  return nil
}

func HandleErrorStatus(res *http.Response, message map[int]string) error {
  errfmt, ok := message[res.StatusCode]
  if !ok {
    return cli.Exit(fmt.Errorf("Invalid response code from the server: %d", res.StatusCode), 6)
  }

  apierr, err := ParseError(res.Body)
  if err != nil {
    return cli.Exit(fmt.Errorf("Couldn't parse error returned by the server: %w", err), 5)
  }

  return cli.Exit(fmt.Errorf(errfmt, apierr), 7)
}
