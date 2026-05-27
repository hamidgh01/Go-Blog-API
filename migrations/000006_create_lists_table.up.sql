CREATE TABLE IF NOT EXISTS lists (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    description VARCHAR(1000),
    isPrivate BOOLEAN DEFAULT true,
    createdAt TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    modifiedAt TIMESTAMP WITH TIME ZONE,
    userID BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_lists_user_id ON lists (userID);


CREATE TABLE IF NOT EXISTS saved_posts_m2m (
    list_id BIGINT NOT NULL REFERENCES lists(id) ON DELETE CASCADE,
    post_id BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    saved_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,

    PRIMARY KEY (list_id, post_id)
);

-- this index will help with queries that filter or join by post_id alone
CREATE INDEX IF NOT EXISTS idx_saved_posts_m2m_post_id ON saved_posts_m2m (post_id);


CREATE TABLE IF NOT EXISTS users_saved_lists_m2m (
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    list_id BIGINT NOT NULL REFERENCES lists(id) ON DELETE CASCADE,
    saved_at TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,

    PRIMARY KEY (user_id, list_id)
);
