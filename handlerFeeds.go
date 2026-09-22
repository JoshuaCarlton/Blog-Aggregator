package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, _ command) error {
	ctx := context.Background()

	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserID(ctx, feed.UserID)
		if err != nil {
			return err
		}

		fmt.Printf("%v, %v, %v\n", feed.Name, feed.Url, user.Name)
	}
	return nil
}
