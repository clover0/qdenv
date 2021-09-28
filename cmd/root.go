package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := buildCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func buildCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "qdenv",
		Short: "Qdenv is a tool to quickly create a programming environment",
		Long:  ``,
	}

	buildInitCmd(rootCmd)
	buildEnterCmd(rootCmd)
	buildHaltCmd(rootCmd)

	return rootCmd
}
