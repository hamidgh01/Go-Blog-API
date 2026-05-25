CREATE TABLE IF NOT EXISTS tags (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(32) NOT NULL UNIQUE -- automatically indexed
);

CREATE TABLE IF NOT EXISTS posts_tags_m2m (
    tag_id BIGINT NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    post_id BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,

    PRIMARY KEY (tag_id, post_id)
);

-- this index will help with queries that filter or join by post_id alone
CREATE INDEX IF NOT EXISTS idx_posts_tags_m2m_postID ON posts_tags_m2m (post_id);
