CREATE TABLE IF NOT EXISTS post_likes_m2m (
    post_id BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    liked_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,

    PRIMARY KEY (post_id, user_id)
);


CREATE TABLE IF NOT EXISTS follows_m2m (
    followed_by BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    followed BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    followed_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,

    PRIMARY KEY (followed_by, followed)
);

-- this index will help with queries that filter or join by followed alone (e.g. fetching followers of a user)
CREATE INDEX IF NOT EXISTS idx_follows_m2m_followed ON follows_m2m (followed);
