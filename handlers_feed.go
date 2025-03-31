package main

import (
	"aggregator/internal/rss"
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feedURL := "https://www.wagslane.dev/index.xml"
	ctx := context.Background()

	feed, err := rss.FetchFeed(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("error fetching RSS feed", err)
	}

	fmt.Println(feed)
	return nil
}
