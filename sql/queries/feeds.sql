-- name: CreateFeed :one
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

-- name: GetFeed :one
SELECT * FROM feeds
WHERE name = $1;

-- name: GetFeedByUrl :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeeds :many 
SELECT feeds.*, users.name as userName FROM feeds, users
WHERE feeds.user_id = users.id;

-- name: DeleteAllFeeds :exec
DELETE FROM feeds;

