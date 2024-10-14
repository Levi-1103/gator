package main

import (
	"context"
	"fmt"
)

func handlerGetUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("command doesn't have any arguments")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't return users")
	}
	loggedInUser := s.config.CurrentUserName

	for _, user := range users {
		if user.Name == loggedInUser {
			fmt.Println("*", user.Name, "(current)")
		} else {
			fmt.Println("*", user.Name)
		}

	}
	return nil
}
