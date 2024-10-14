package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	req.Header.Add("User-Agent", "gator")
	if err != nil {
		return &RSSFeed{}, err
	}

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	defer res.Body.Close()

	var rssFeed RSSFeed

	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return &RSSFeed{}, err
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	items := rssFeed.Channel.Item

	for i := range items {
		item := &items[i]
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
	}

	return &rssFeed, nil
}
