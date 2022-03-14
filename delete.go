package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

func Delete(c *cli.Context) error {
  bp, err := LoadBlogPost(".blogpost.json")
  if err != nil {
    if errors.Is(err, os.ErrNotExist) {
      return cli.Exit(fmt.Errorf("The entry hasn't been posted. Did you mean `post`?"), 10)
    }
    return cli.Exit(fmt.Errorf("Couldn't load data about the blog post: %w", err), 11)
  }

  path, err := APIPath("blog", bp.Meta.Slug)
  if err != nil {
    return cli.Exit(fmt.Errorf("Couldn't generate the api route: %w", err), 13)
  }

  req, _ := http.NewRequest(http.MethodDelete, path, nil) // path is guaranteed to be parsable, so the error can be ignored
  res, err := (&http.Client{}).Do(req)
  if err != nil {
    return cli.Exit(fmt.Errorf("Couldn't complete request: %w", err), 12)
  }
  defer res.Body.Close()

  if res.StatusCode != http.StatusOK {
    return HandleErrorStatus(res, map[int]string{
      http.StatusBadRequest: "There's something wrong with this client",
      http.StatusNotFound: "The blog doesn't seem to exist, was it deleted by some other client?",
      http.StatusInternalServerError: "Server Error: %w",
    })
  }

  os.Remove(".blogpost.json")

  fmt.Println("The entry was successfully deleted")
  return nil
}
