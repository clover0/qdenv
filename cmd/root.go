package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "qdenv",
	Short: "qdenv is a tool for building programming environment quickly",
	Long:  ``,
}

func init() {
	buildInitCmd()
	buildEnterCmd()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
