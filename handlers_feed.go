package main

import (
	"aggregator/internal/database"
	"aggregator/internal/rss"
	"context"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"time"
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

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Arguments) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	ctx := context.Background()
	feedName := cmd.Arguments[0]
	feedURL := cmd.Arguments[1]

	currentTime := time.Now()
	newFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	})

	if err != nil {
		return fmt.Errorf("error creating feed in database", err)
	}

	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    user.ID,
		FeedID:    newFeed.ID,
	})

	fmt.Println(newFeed)

	return nil
}
func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	ctx := context.Background()

	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("error retrieving feeds in database", err)
	}

	for _, feed := range feeds {
		fmt.Println("Feeds")
		fmt.Printf("Name: %s | URL: %s | Created by: %s\n", feed.Name, feed.Url, feed.Username)
	}
	return nil
}
