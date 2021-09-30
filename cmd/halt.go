package cmd

import (
	"github.com/spf13/cobra"

	"github.com/clover0/qdenv/util"
)

func buildHaltCmd(rootCmd *cobra.Command) {
	var cmdEnter = &cobra.Command{
		Use:   "halt",
		Short: "Halt environment",
		Long:  ``,
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runHalt()
		},
	}

	rootCmd.AddCommand(cmdEnter)
}

func runHalt() error {
	return util.Execw("docker-compose", []string{"stop"})
}
