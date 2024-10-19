package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Levi-1103/gator/internal/database"
	"github.com/google/uuid"
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

		pubTime, _ := time.Parse(time.RFC1123Z, item.PubDate)
		publishedAt := sql.NullTime{
			Time:  pubTime,
			Valid: true,
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: now,
			UpdatedAt: now,
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: publishedAt,
			FeedID:      nextFeed.ID,
		},
		)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}

	fmt.Println("Collected Posts")

	return nil
}
