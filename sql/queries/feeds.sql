-- name: AddFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT
    name,
    url,
    user_id
FROM
    feeds;

-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES(
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT
    inserted_feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follows
INNER JOIN feeds ON inserted_feed_follows.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follows.user_id = users.id;

-- name: GetFeed :one
SELECT
    *
FROM
    feeds
WHERE
    url = $1
LIMIT 1;

-- name: GetFeedFollowsForUser :many
SELECT
    feeds.name AS feed_name,
    users.name AS user_name
FROM
    feed_follows
    INNER JOIN users ON feed_follows.user_id = users.id
    INNER JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE
    users.name = $1;
