package main

import (
	"context"
	"fmt"

	"github.com/Levi-1103/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("command has no arguments")
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Println("Feed Follows for " + user.Name)
	for _, feed := range feeds {
		fmt.Println("*", feed.FeedName)
	}

	return nil
}
