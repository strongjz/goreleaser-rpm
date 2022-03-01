package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"sigs.k8s.io/release-utils/version"
)

func main() {
	var hello = &cobra.Command{
		Use:   "hello",
		Short: "prints out hello",
		Long:  `prints out hello`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello World")
		},
	}

	var name = &cobra.Command{
		Use:   "name",
		Short: "prints out hello name",
		Long:  `prints out hello name`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello %v\n", args[0])
		},
	}

	hello.AddCommand(version.Version())
	hello.AddCommand(name)
	hello.Execute()
}
