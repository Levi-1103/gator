package main

import (
	"fmt"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("command requires arguments: register <url>")
	}

	return nil
}
