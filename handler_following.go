package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("command has no arguments")
	}

	current_user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), current_user.ID)
	if err != nil {
		return err
	}

	fmt.Println("Feed Follows for " + current_user.Name)
	for _, feed := range feeds {
		fmt.Println("*", feed.FeedName)
	}

	return nil
}
