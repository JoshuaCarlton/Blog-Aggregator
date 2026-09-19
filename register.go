package main

import (
	"context"
	"fmt"
	"time"

	"github.com/JoshuaCarlton/Blog-Aggregator/internal/config"
	"github.com/JoshuaCarlton/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("must provide a username to register")
	}

	ctx := context.Background()

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
	}

	user, err := s.db.CreateUser(ctx, userParams)
	if err != nil {
		return err
	}
	fmt.Printf("user registered with id: %v, created at: %v, updated at %v, name: %v\n", user.ID, user.CreatedAt, user.UpdatedAt, user.Name)
	err = config.SetUser(s.config, cmd.arguments[0])
	if err != nil {
		return err
	}

	return nil
}
