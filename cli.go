package main

import (
  "github.com/urfave/cli/v2"
)

func MakeApp() *cli.App {
  return &cli.App{
    Name: "blog",
    Usage: "a cli tool to manage blog entries",
    Description: "a cli tool to post/put/delete blog entries by interfacing with the api running at $BLOG_API_URL",
    Commands: []*cli.Command{
      {
        Name: "init",
        Aliases: []string{ "i" },
        Description: "initializes the current directory as a blog post entry",
        Usage: "initializes the current directory",
        Action: Init,
      },
      {
        Name: "post",
        Aliases: []string{ "po" },
        Description: "posts the contents of the current blog directory to $BLOG_API_URL",
        Usage: "posts a blog entry",
        Action: Post,
      },
      {
        Name: "put",
        Aliases: []string{ "pu" },
        Description: "puts the contents of the current blog directory to $BLOG_API_URL/:slug",
        Usage: "puts a blog entry",
        Action: Put,
      },
      {
        Name: "delete",
        Aliases: []string{ "d" },
        Description: "deletes the contents of the current blog directory from $BLOG_API_URL/:slug",
        Usage: "deletes a blog entry",
        Action: Delete,
      },
    },
  }
}
