package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/JoshuaCarlton/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.arguments) > 0 {
		num, err := strconv.ParseInt(cmd.arguments[0], 10, 32)
		if err != nil {
			return err
		}
		limit = int32(num)
	}

	ctx := context.Background()

	brwoseParams := database.GetPostsForUserParams{UserID: user.ID, Limit: limit}

	posts, err := s.db.GetPostsForUser(ctx, brwoseParams)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Println(post.Url)
	}
	return nil
}
