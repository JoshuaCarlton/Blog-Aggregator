package main

import (
	"context"
	"fmt"

	"github.com/JoshuaCarlton/Blog-Aggregator/internal/config"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("login requires a username argument\n")
	}

	ctx := context.Background()

	_, err := s.db.GetUser(ctx, cmd.arguments[0])
	if err != nil {
		return err
	}

	err = config.SetUser(s.config, cmd.arguments[0])
	if err != nil {
		return err
	}
	fmt.Printf("username has been set to %s\n", cmd.arguments[0])
	return nil
}
