-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT feeds.name, posts.* FROM posts, feeds, users
WHERE users.id = $1 and
users.id = feeds.user_id
and feeds.id = posts.feed_id
ORDER BY posts.published_at DESC
LIMIT $2;




