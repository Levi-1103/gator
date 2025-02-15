-- name: GetFeeds :many
SELECT feeds.name, feeds.url, users.name FROM feeds INNER JOIN users ON user_id = users.id;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;