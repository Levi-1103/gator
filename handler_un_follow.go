package main

import (
	"context"
	"fmt"

	"github.com/Levi-1103/gator/internal/database"
)

func handlerUnFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("command requires arguments: unfollow <url>")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.db.UnFollowFeed(context.Background(), database.UnFollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return fmt.Errorf("couldn't delete feed follow: %w", err)
	}

	fmt.Printf("%s unfollowed successfully!\n", feed.Name)

	return nil
}
