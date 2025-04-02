package main

import (
	"aggregator/internal/database"
	"context"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"time"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <feed URL>", cmd.Name)
	}

	feedURL := cmd.Arguments[0]
	ctx := context.Background()

	feed, err := s.db.GetFeedByUrl(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("error retrieving feed from DB: %w", err)
	}

	currentTime := time.Now()

	newFollow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("error inserting new feed follow: %w", err)
	}

	fmt.Printf("New follow for feed %s by user %s\n", newFollow.Feedname, newFollow.Username)
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <feed URL>", cmd.Name)
	}

	feedURL := cmd.Arguments[0]
	ctx := context.Background()

	feed, err := s.db.GetFeedByUrl(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("error retrieving feed by URL: %w", err)
	}

	err = s.db.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return fmt.Errorf("error removing feed from DB: %w", err)
	}
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	ctx := context.Background()

	feeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("error retrieving feeds from user: %w", err)
	}
	fmt.Printf("Feeds followed by user %s\n", user.Name)
	for _, feed := range feeds {
		fmt.Println(feed)
	}

	return nil
}
