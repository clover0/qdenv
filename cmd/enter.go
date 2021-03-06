package cmd

import (
	"github.com/spf13/cobra"

	"github.com/clover0/qdenv/util"
)

func buildEnterCmd(rootCmd *cobra.Command) {
	var cmdEnter = &cobra.Command{
		Use:   "enter",
		Short: "Enter environment",
		Long: `Enter environment. 
If not started environment(docker container), start container before enter environment`,
		Args: cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runEnter()
		},
	}

	rootCmd.AddCommand(cmdEnter)
}

func runEnter() error {
	if err := util.Execw("docker-compose", []string{"up", "-d"}); err != nil {
		return err
	}

	return util.Execw("docker-compose", []string{"exec", "qdenv", "sh", "-c", "bash || sh"})
}
