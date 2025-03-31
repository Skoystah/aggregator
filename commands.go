package main

import (
	"fmt"
)

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	cliCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cliCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, exists := c.cliCommands[cmd.Name]
	if !exists {
		return fmt.Errorf("command does not exists")
	}

	return f(s, cmd)
}
