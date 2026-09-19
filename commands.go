package main

import "fmt"

type command struct {
	handlerName string
	arguments   []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.handlerName]
	if !ok {
		return fmt.Errorf("invalid command")
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func newCommands() commands {
	var cm map[string]func(*state, command) error = map[string]func(*state, command) error{}
	cmds := commands{cm}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	return cmds
}
