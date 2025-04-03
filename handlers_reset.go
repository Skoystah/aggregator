package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	ctx := context.Background()

	err := s.db.DeleteAllUsers(ctx)
	if err != nil {
		return fmt.Errorf("error resetting database: %w", err)
	}
	fmt.Println("Reset user database successful")
	return nil
}
