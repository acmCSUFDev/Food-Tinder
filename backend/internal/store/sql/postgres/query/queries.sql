-- User gets any user's public data.
-- name: User :one
SELECT
    *
FROM
    public_users
WHERE
    username = $1
LIMIT 1;

-- Self gets the current user's private data.
-- name: Self :one
SELECT
    *
FROM
    private_users
WHERE
    username = $1
LIMIT 1;

-- UserPasshash gets a user's password hash.
-- name: UserPasshash :one
SELECT
    passhash
FROM
    users
WHERE
    username = $1
LIMIT 1;

-- name: CreateUser :exec
INSERT INTO users (username, passhash)
    VALUES ($1, $2);

-- name: UpdateUser :exec
UPDATE
    users
SET
    display_name = $2,
    avatar = $3,
    bio = $4,
    birthday = $5,
    food_preferences = $6
WHERE
    username = $1;

-- name: ChangePassword :exec
UPDATE
    users
SET
    passhash = $2
WHERE
    username = $1;

-- name: CreateSession :one
INSERT INTO sessions (token, username, expiry, metadata)
    VALUES ($1, $2, new_session_expiry (), $3)
RETURNING
    *;

-- ValidateSession validates the session and returns the username. An error is
-- returned if the session doesn't exist anymore.
-- name: ValidateSession :one
UPDATE
    sessions
SET
    expiry = new_session_expiry ()
WHERE
    token = $1
    AND expiry < now()
RETURNING
    username,
    expiry,
    metadata;

-- TokenExists returns 1 if the token exists or 0.
-- name: TokenExists :one
SELECT
    COUNT(*)
FROM
    sessions
WHERE
    token = $1;

-- TODO: consider returning an indication
-- name: DeleteSession :exec
DELETE FROM sessions
WHERE token = $1;

-- NextPosts paginates the list of posts.
-- name: NextPosts :many
SELECT
    posts.*,
    count_likes (posts.id) AS likes
FROM
    posts
WHERE
    id < $1
    OR $1 = 0
LIMIT 10;

-- LikedPosts returns a user's liked posts.
-- name: LikedPosts :many
SELECT
    posts.*,
    count_likes (posts.id) AS likes
FROM
    posts
WHERE
    posts.id IN (
        SELECT
            liked_posts.post_id
        FROM
            liked_posts
        WHERE
            liked_posts.username = $1);

-- PostLikeCount returns the like count of the post with the given ID.
-- name: PostLikeCount :one
SELECT
    COUNT(*)
FROM
    liked_posts
WHERE
    post_id = $1;

-- DeletePost deletes a post.
-- name: DeletePost :execrows
DELETE FROM posts
WHERE id = $1
    AND username = $2;

-- CreatePost creates a new post.
-- name: CreatePost :exec
INSERT INTO posts (id, username, cover_hash, images, description, tags, location)
    VALUES ($1, $2, $3, $4, $5, $6, $7);

