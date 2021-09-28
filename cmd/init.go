package cmd

import (
	"os"
	"strings"
	gotmpl "text/template"

	"github.com/spf13/cobra"

	"qdenv/template"
	"qdenv/util"
)

type initArgs struct {
	raw       []string
	ImageName string
	TagName   string
}

func NewInitArgs(args []string) initArgs {
	intn := strings.Split(args[0], ":")
	return initArgs{raw: args, ImageName: intn[0], TagName: intn[1]}
}

func buildInitCmd() {
	var cmdInit = &cobra.Command{
		Use:   "init [image-name:tag-name]",
		Short: "Init environment",
		Long: `Initialize environment by target.
				e.g) init python:3.9`,
		Args: cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			a := NewInitArgs(args)
			if err := runInit(a); err != nil {
				return err
			}
			return nil
		},
	}

	rootCmd.AddCommand(cmdInit)
}

func runInit(args initArgs) (err error) {
	if err = createFile("Dockerfile", template.DockerfileTmpl, args); err != nil {
		return
	}

	if err = createFile("docker-compose.yml", template.DockerComposeTmpl, args); err != nil {
		return
	}

	if err = buildImage(); err != nil {
		return err
	}
	return
}

func createFile(name string, tmpl string, args initArgs) error {
	t, err := gotmpl.New(name).Parse(tmpl)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	if err = t.Execute(f, args); err != nil {
		return err
	}
	return nil
}

func buildImage() error {
	return util.Execw("docker-compose", []string{"build"})
}
