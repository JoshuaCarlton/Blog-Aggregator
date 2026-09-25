package main

import (
	"context"
	"fmt"
	"time"

	"github.com/JoshuaCarlton/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("you must provide the name and url of the feed")
	}

	ctx := context.Background()

	feedParams := database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
		Url:       cmd.arguments[1],
		UserID:    user.ID,
	}

	feed, err := s.db.AddFeed(ctx, feedParams)
	if err != nil {
		return err
	}

	cmd.arguments[0] = feed.Url
	err = handlerFollow(s, cmd, user)
	return err
}
