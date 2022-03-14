package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

func Put(c *cli.Context) error {
  bp, err := LoadBlogPost(".blogpost.json")
  if err != nil {
    if errors.Is(err, os.ErrNotExist) {
      return cli.Exit(fmt.Errorf("The entry hasn't been posted. Did you mean `post`?"), 12)
    }
    return cli.Exit(fmt.Errorf("Couldn't load data about the blog post: %w", err), 13)
  }

  post, err := NewCreateParams()
  if err != nil {
    return cli.Exit(fmt.Errorf("Couldn't load params: %w", err), 14)
  }

  bp.Meta.Title = post.Title
  bp.Meta.Summary = post.Summary
  bp.Meta.Tags = post.Tags
  bp.Content = post.Content

  path, err := APIPath("blog", bp.Meta.Slug)
  if err != nil {
    return cli.Exit(fmt.Errorf("Couldn't generate api route: %w", err), 15)
  }

  res, err := PutJson(path, bp)
  if err != nil {
    return cli.Exit(fmt.Errorf("Couldn't complete request: %w", err), 16)
  }
  defer res.Body.Close()

  if res.StatusCode != http.StatusOK {
    return HandleErrorStatus(res, map[int]string{
      http.StatusBadRequest: "There's something wrong with this client",
      http.StatusNotFound: "The blog post doesn't seem to exist, was it deleted by some other client?",
      http.StatusInternalServerError: "Server Error: %w",
    })
  }

  err = json.NewDecoder(res.Body).Decode(&bp)
  if err != nil {
    return cli.Exit(fmt.Errorf("Couldn't parse the response returned by the server: %w", err), 17)
  }

  err = bp.SaveBlogPost(".blogpost.json")
  if err != nil {
    return cli.Exit(fmt.Errorf("Couldn't save the blog post data: %w", err), 18)
  }

  fmt.Println("Entry was succesfully put at", "/" + bp.Meta.Slug)
  return nil
}
