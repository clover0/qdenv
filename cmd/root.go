package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qdenv",
	Short: "Qdenv is a tool to quickly create a programming environment",
	Long:  ``,
}

func init() {
	buildInitCmd()
	buildEnterCmd()
	buildHaltCmd()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
