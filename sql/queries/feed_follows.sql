-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING (select feeds.name from feeds where feeds.id = feed_id) as feedName,
(select users.name from users where id = $4) as userName;

-- name: GetFeedFollowsForUser :many
SELECT feeds.name FROM feed_follows, feeds
WHERE feed_follows.user_id = $1 AND
feed_follows.feed_id = feeds.id;


