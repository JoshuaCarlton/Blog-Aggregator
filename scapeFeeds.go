package main

import (
	"context"
	"fmt"
	"time"

	"github.com/JoshuaCarlton/Blog-Aggregator/internal/database"
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

	fmt.Println(rssfeed.Channel.Title)
	for _, item := range rssfeed.Channel.Items {
		fmt.Println(item.Title)
	}

	return nil
}
