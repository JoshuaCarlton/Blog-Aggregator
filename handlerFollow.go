package main

import (
	"context"
	"fmt"
	"time"

	"github.com/JoshuaCarlton/Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("you must provide the url of the feed you want to follow")
	}

	ctx := context.Background()

	feed, err := s.db.GetFeed(ctx, cmd.arguments[0])
	if err != nil {
		return err
	}

	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	follow, err := s.db.CreateFeedFollow(ctx, followParams)
	if err != nil {
		return err
	}

	fmt.Printf("%v has been followed by %v\n", follow.FeedName, follow.UserName)
	return nil
}
