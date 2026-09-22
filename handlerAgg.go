package main

import (
	"context"
	"fmt"
)

func handlerAgg(_ *state, _ command) error {
	ctx := context.Background()

	url := "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(ctx, url)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
