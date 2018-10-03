package main

import (
    "os"
    "fmt"
    "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()

  app.Name = "eb-ssh"
  app.Usage = "This app echo input arguments"
  app.Version = "0.0.1"

  app.Action = func (context *cli.Context) error {
    if context.String("env-name") != "" {
      fmt.Println("env-name: " + context.Args().Get(0))
    } else if context.String("region") != "" {
      fmt.Println("region: " + context.Args().Get(0))
    } else {
      fmt.Println(context.Args().Get(0))
    }
    return nil
  }

  app.Flags = []cli.Flag{
    cli.StringFlag {
      Name: "env-name, e",
      Usage: "",
    },
    cli.StringFlag {
      Name: "region, r",
      Usage: "",
    },
    cli.StringFlag {
      Name: "profile, p",
      Usage: "",
    },
    cli.StringFlag {
      Name: "keyfile, k",
      Usage: "",
    },
    cli.StringFlag {
      Name: "proxy-host, P",
      Usage: "",
    },
  }

  app.Run(os.Args)
}
