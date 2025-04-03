package main

import (
	"aggregator/internal/config"
	"aggregator/internal/database"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.URL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	currentState := &state{cfg: &cfg, db: dbQueries}

	commands := commands{cliCommands: make(map[string]func(*state, command) error)}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	commands.register("browse", middlewareLoggedIn(handlerBrowse))

	iptArguments := os.Args
	if len(iptArguments) < 2 {
		log.Fatal("error - too few arguments")
		return
	}

	cmdName := iptArguments[1]
	cmdArgs := iptArguments[2:]
	cmd := command{Name: cmdName, Arguments: cmdArgs}

	err = commands.run(currentState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		ctx := context.Background()

		user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("error retrieving current user: %w", err)
		}
		return handler(s, cmd, user)
	}
}
