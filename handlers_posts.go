package main

import (
	"aggregator/internal/database"
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <post limit>", cmd.Name)
	}

	ctx := context.Background()
	postLimit, err := strconv.Atoi(cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("error converting str to int %w", err)
	}

	posts, err := s.db.GetPostsForUser(ctx, database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: int32(postLimit)})

	if err != nil {
		return fmt.Errorf("error retrieving posts in database: %w", err)
	}

	fmt.Println("Posts:")
	for _, post := range posts {
		fmt.Printf("Feed: %s\n", post.FeedID)
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Description: %s\n", post.Description.String)
		fmt.Printf("Published at: %s\n\n", post.CreatedAt)
	}
	return nil
}
