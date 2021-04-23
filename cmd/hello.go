package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Hello(name string) string {
	return "hello " + name
}

var helloCmd = &cobra.Command{
	Use: "hello",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(Hello(args[0]))
		}
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
