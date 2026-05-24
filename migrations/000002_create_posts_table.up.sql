CREATE TYPE PostStatus AS ENUM ('draft', 'published', 'rejected', 'deleted-by-author');

CREATE TABLE IF NOT EXISTS posts (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    content TEXT,
    status PostStatus DEFAULT 'draft',
    isPrivate BOOLEAN DEFAULT false,
    createdAt TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    modifiedAt TIMESTAMP WITH TIME ZONE,
    firstPublishedAt TIMESTAMP WITH TIME ZONE,
    userID BIGINT NOT NULL DEFAULT 0 REFERENCES users(id) ON DELETE SET DEFAULT
);

CREATE INDEX IF NOT EXISTS idx_posts_title ON posts (title);
CREATE INDEX IF NOT EXISTS composite_idx_posts_userID_pubAt ON posts (userID, firstPublishedAt);
