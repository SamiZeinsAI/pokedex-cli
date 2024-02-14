package main

import "os"

func exit_cmd(cfg *config, args ...string) error {
	os.Exit(1)
	return nil
}
