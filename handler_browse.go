package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Levi-1103/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) != 1 && len(cmd.args) != 0 {
		return fmt.Errorf("command requires arguments: browse optional:<limit>")
	}

	if len(cmd.args) == 1 {
		specifiedLimit, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalid limit input")
		}
		limit = specifiedLimit
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}

	for _, post := range posts {
		fmt.Println("*", post.Title)
		fmt.Println()
		fmt.Println("*", post.Description)
		fmt.Println()
		fmt.Println("*", post.Url)
		fmt.Println()
		fmt.Println("*", post.PublishedAt)
		fmt.Println("==================")
	}

	return nil
}
