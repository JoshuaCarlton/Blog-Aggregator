package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, _ command) error {
	ctx := context.Background()

	user, err := s.db.GetUser(ctx, s.config.CurrentUserName)
	if err != nil {
		return err
	}

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
