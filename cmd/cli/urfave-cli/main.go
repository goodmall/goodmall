package main

import (
  "os"
  "sort"
  "fmt"

  "github.com/urfave/cli"

  "github.com/goodmall/goodmall/pods/user/usecase"
)

func main() {
  app := cli.NewApp()

  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "lang, l",
      Value: "english",
      Usage: "Language for the greeting",
    },
    cli.StringFlag{
      Name: "userhelp, uh",
      Usage: "get the help from module user",
    },
  }

  app.Commands = []cli.Command{
    {
      Name:    "userhelp",
      Aliases: []string{"uh"},
      Usage:   "how to use the user module",
      Action:  func(c *cli.Context) error {
		 
		// userInteractor := usecase.NewUserInteractor()
		userInteractor :=  usecase.NewUsecase(/* 依赖暂缺 */).NewUserInteractor() // usecase.NewUserInteractor()
		response := userInteractor.Help() 
		fmt.Println(response)


        return nil
      },
    },
  }

  sort.Sort(cli.FlagsByName(app.Flags))
  sort.Sort(cli.CommandsByName(app.Commands))

  app.Run(os.Args)
}