package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/JoshuaCarlton/gator/internal/database"
	"github.com/google/uuid"
)

func savePost(s *state, item RSSItem, feedId uuid.UUID) {
	ctx := context.Background()

	description := sql.NullString{String: "", Valid: false}
	if item.Description != "" {
		description.String = item.Description
		description.Valid = true
	}

	publishedAt := sql.NullTime{Time: time.Time{}, Valid: false}
	parseTime, err := time.Parse(time.RFC1123Z, item.PubDate)
	if err != nil {
		fmt.Println("couldnt parse pubdate", err.Error())
	} else {
		publishedAt.Time = parseTime
		publishedAt.Valid = true
	}

	postParams := database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       item.Title,
		Url:         item.Link,
		Description: description,
		PublishedAt: publishedAt,
		FeedID:      feedId,
	}

	_, err = s.db.CreatePost(ctx, postParams)
	if err != nil {
		if err.Error() != `pq: duplicate key value violates unique constraint "posts_url_key" (23505)` {
			fmt.Println(err.Error())
		}
	}
}
