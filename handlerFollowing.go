package main

import (
	"context"
	"fmt"

	"github.com/JoshuaCarlton/gator/internal/database"
)

func handlerFollowing(s *state, _ command, user database.User) error {
	ctx := context.Background()

	following, err := s.db.GetFeedFollowsForUser(ctx, user.Name)
	if err != nil {
		return err
	}

	fmt.Printf("user %v is following:\n", user.Name)
	for _, follow := range following {
		fmt.Println(follow.FeedName)
	}
	return nil
}
