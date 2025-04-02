package main

import (
	"aggregator/internal/config"
	"aggregator/internal/database"
	"database/sql"
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

	//TODO - const for login name?
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", handlerAddFeed)
	commands.register("feeds", handlerFeeds)
	commands.register("follow", handlerFollow)
	commands.register("following", handlerFollowing)

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
