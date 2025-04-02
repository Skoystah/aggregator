package main

import (
	"aggregator/internal/database"
	"context"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"time"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	userName := cmd.Arguments[0]
	ctx := context.Background()

	_, err := s.db.GetUser(ctx, userName)
	if err != nil {
		return fmt.Errorf("error retrieving user in database", err)
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("error setting user name: %w", err)
	}

	fmt.Println("Username has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	userName := cmd.Arguments[0]

	ctx := context.Background()

	currentTime := time.Now()
	registeredUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Name:      userName,
	})

	if err != nil {
		return fmt.Errorf("error creating user in database", err)
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("error setting user name: %w", err)
	}

	fmt.Println("User registered in database: %w", registeredUser)
	return nil
}

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

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Arguments) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	ctx := context.Background()

	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("error retrieving users in database", err)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
