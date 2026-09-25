package main

import (
	"context"
	"fmt"
	"time"

	"github.com/JoshuaCarlton/gator/internal/database"
)

func scrapeFeeds(s *state) error {
	ctx := context.Background()

	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{ID: feed.ID, UpdatedAt: time.Now()})
	if err != nil {
		return err
	}

	rssfeed, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}
	count := 0
	fmt.Println(rssfeed.Channel.Title)
	for _, item := range rssfeed.Channel.Items {
		savePost(s, item, feed.ID)
		count += 1
	}
	fmt.Printf("saved %d posts\n", count)
	return nil
}
