-- +goose Up
CREATE TABLE users (
    username varchar(50) PRIMARY KEY,
    passhash bytea NOT NULL,
    display_name text,
    avatar text,
    bio text,
    birthday date,
    food_preferences json NOT NULL
);

CREATE VIEW public_users AS
SELECT
    username,
    display_name,
    avatar,
    bio
FROM
    users;

CREATE VIEW private_users AS
SELECT
    username,
    display_name,
    avatar,
    bio,
    birthday,
    food_preferences
FROM
    users;

CREATE TABLE sessions (
    token text PRIMARY KEY,
    username varchar(50) NOT NULL REFERENCES users (username) ON DELETE CASCADE ON UPDATE CASCADE,
    expiry timestamp NOT NULL,
    metadata json NOT NULL
);

CREATE TABLE posts (
    id bigint PRIMARY KEY,
    username varchar(50) NOT NULL REFERENCES users (username) ON DELETE CASCADE ON UPDATE CASCADE,
    cover_hash text,
    images text[],
    description text,
    tags text[],
    location text
);

CREATE TABLE liked_posts (
    username varchar(50) NOT NULL REFERENCES users (username) ON DELETE CASCADE ON UPDATE CASCADE,
    post_id bigint NOT NULL REFERENCES posts (id) ON DELETE CASCADE,
    liked_at timestamp NOT NULL,
    UNIQUE (username, post_id)
);

CREATE FUNCTION new_session_expiry (timestamp DEFAULT now())
    RETURNS timestamp
    AS 'SELECT $1 + (7 * interval ''1 day'')'
    LANGUAGE SQL;

CREATE FUNCTION count_likes (post_id bigint)
    RETURNS bigint
    AS 'SELECT COUNT(*) FROM liked_posts WHERE post_id = $1'
    LANGUAGE SQL;

-- +goose Down
