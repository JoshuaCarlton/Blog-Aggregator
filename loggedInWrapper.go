package main

import (
	"context"

	"github.com/JoshuaCarlton/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	innerFunc := func(s *state, cmd command) error {
		ctx := context.Background()

		user, err := s.db.GetUser(ctx, s.config.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
	return innerFunc
}
