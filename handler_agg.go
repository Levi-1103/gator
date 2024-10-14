package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("command has no arguments")
	}

	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed request: %w", err)
	}

	fmt.Println(feed)

	return nil
}
