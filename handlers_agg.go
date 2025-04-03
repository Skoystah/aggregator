package main

import (
	"aggregator/internal/database"
	"aggregator/internal/rss"
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <time between req>", cmd.Name)
	}

	time_between_reqs, err := time.ParseDuration(cmd.Arguments[0])

	if err != nil {
		return fmt.Errorf("error parsing time: %w", err)
	}

	fmt.Printf("Collecting feeds every %s", time_between_reqs)

	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("error scraping feed: %w", err)
		}
	}

}

func scrapeFeeds(s *state) error {

	ctx := context.Background()

	nextFeed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("error retrieving feed to fetch: %w", err)
	}

	err = s.db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
		ID:        nextFeed.ID,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	feed, err := rss.FetchFeed(ctx, nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching RSS feed: %w", err)
	}

	for _, item := range feed.Channel.Item {
		//		fmt.Printf("Adding post %v\n\n", item)
		currentTime := time.Now()

		postDescription := sql.NullString{String: item.Description}
		if item.Description > "" {
			postDescription.Valid = true
		} else {
			postDescription.Valid = false
		}

		pubTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return fmt.Errorf("Error formatting pubdate %s: %w", item.PubDate, err)
		}

		_, err = s.db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
			Title:       item.Title,
			Url:         item.Link,
			Description: postDescription,
			PublishedAt: pubTime,
			FeedID:      nextFeed.ID,
		})
		if err != nil {
			return fmt.Errorf("Error creating post %v: %w", item, err)
		}
	}
	return nil
}
