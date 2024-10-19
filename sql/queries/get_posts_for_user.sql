-- name: GetPostsForUser :many
SELECT posts.*
FROM posts
INNER JOIN feeds ON posts.feed_id = feeds.id
INNER JOIN feed_follows ON feed_follows.feed_id = feeds.id
WHERE feed_follows.user_id = $1
ORDER BY posts.updated_at DESC
LIMIT $2;