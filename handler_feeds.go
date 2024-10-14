package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("command has no arguments")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error couldn't get feeds: %w", err)
	}

	fmt.Println("RSS Feeds")
	fmt.Println()
	for _, feed := range feeds {
		fmt.Println("Feed Name:", feed.Name)
		fmt.Println("Feed Url:", feed.Url)
		fmt.Println("Created By:", feed.Name_2)
		fmt.Println("-------")
	}

	return nil
}
