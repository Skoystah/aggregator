package main

import (
	"aggregator/internal/database"
	"context"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"time"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <feed URL>", cmd.Name)
	}

	feedURL := cmd.Arguments[0]
	ctx := context.Background()

	currentUser, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error retrieving current user: %w", err)
	}

	feed, err := s.db.GetFeedByUrl(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("error retrieving feed from DB: %w", err)
	}

	currentTime := time.Now()

	newFollow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("error inserting new feed follow: %w", err)
	}

	fmt.Printf("New follow for feed %s by user %s\n", newFollow.Feedname, newFollow.Username)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	ctx := context.Background()

	currentUser, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error retrieving current user: %w", err)
	}

	feeds, err := s.db.GetFeedFollowsForUser(ctx, currentUser.ID)
	if err != nil {
		return fmt.Errorf("error retrieving feeds from user: %w", err)
	}
	fmt.Printf("Feeds followed by user %s\n", s.cfg.CurrentUserName)
	for _, feed := range feeds {
		fmt.Println(feed)
	}

	return nil
}
