package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Levi-1103/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("command requires argument: agg <time_between_reqs> example: 1h")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("not a valid time string")
	}

	fmt.Println("Collecting feeds every: ", cmd.args[0])

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		go scrapeFeeds(s)
	}

}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	nullTime := sql.NullTime{
		Time:  now,
		Valid: true,
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:            nextFeed.ID,
		LastFetchedAt: nullTime,
		UpdatedAt:     now,
	})
	if err != nil {
		return err
	}

	feedData, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	fmt.Println()

	for _, item := range feedData.Channel.Item {
		fmt.Println(item.Title)
	}

	fmt.Println()

	return nil
}
