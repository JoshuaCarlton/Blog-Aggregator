package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/JoshuaCarlton/Blog-Aggregator/internal/config"
	"github.com/JoshuaCarlton/Blog-Aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	cmds := newCommands()

	args := os.Args

	if len(args) < 2 {
		fmt.Println("you must provide a command")
		os.Exit(1)
	}
	cmd := command{
		handlerName: args[1],
		arguments:   args[2:],
	}
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	dbQueries := database.New(db)

	state := state{
		config: &cfg,
		db:     dbQueries,
	}

	err = cmds.run(&state, cmd)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
