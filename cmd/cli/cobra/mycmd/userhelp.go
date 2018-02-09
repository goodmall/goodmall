package mycmd

import (
  "fmt"

  "github.com/spf13/cobra"

  "github.com/goodmall/goodmall/pods/user/usecase"
)

func init() {
  rootCmd.AddCommand(userhelpCmd)
}

var userhelpCmd = &cobra.Command{
  Use:   "userhelp",
  Short: "Print the help info for the usecase userhelp which is in user module",
  Long:  `All usecase has a help info this is the user-help 's `,
  Run: func(cmd *cobra.Command, args []string) {
	 // fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	 
	 // userInteractor := usecase.NewUserInteractor()
	 userInteractor :=  usecase.NewUsecase(/* 依赖暂缺 */).NewUserInteractor() // usecase.NewUserInteractor()
	 response := userInteractor.Help() 
	 fmt.Println(response)

  },
}