package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("command requires arguments: login <username>")
	}

	username := cmd.args[0]

	err := s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User %s has been set\n", username)

	return nil
}
