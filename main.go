package main

import (
	"fmt"
	"os"

	"github.com/JoshuaCarlton/Blog-Aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	cmds := newCommands()
	state := state{&cfg}
	args := os.Args
	if len(args) < 2 {
		fmt.Println("you must provide a command")
		os.Exit(1)
	}
	cmd := command{
		handlerName: args[1],
		arguments:   args[2:],
	}
	err = cmds.run(&state, cmd)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
