package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Levi-1103/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("command requires arguments: register <username>")
	}

	username := cmd.args[0]

	dbUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(), CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(), Name: username,
	})

	if err != nil {
		return fmt.Errorf("user with name '%s' already exists", username)
	}

	s.config.SetUser(username)

	fmt.Println("User was created")
	fmt.Println(dbUser)

	return nil
}
