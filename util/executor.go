package util

import (
	"os"
	"os/exec"
)

func Execw(cmd string, args []string) error {
	c := exec.Command(cmd, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Start()
	if err != nil {
		return err
	}

	err = c.Wait()
	if err != nil {
		return err
	}

	return nil
}
