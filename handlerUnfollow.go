package main

import (
	"context"
	"fmt"

	"github.com/JoshuaCarlton/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("you must provide the url of the feed you want to unfollow")
	}

	ctx := context.Background()

	feed, err := s.db.GetFeed(ctx, cmd.arguments[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID})
	return err
}
