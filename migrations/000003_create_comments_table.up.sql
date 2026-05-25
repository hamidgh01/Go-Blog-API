CREATE TYPE CommentStatus AS ENUM ('published', 'hidden-by-Admin', 'deleted-by-commenter');

CREATE TABLE IF NOT EXISTS comments (
    id BIGSERIAL PRIMARY KEY,
    content VARCHAR(1500) NOT NULL,
    status CommentStatus DEFAULT 'published',
    createdAt TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp,
    modifiedAt TIMESTAMP WITH TIME ZONE,
    userID BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    postParentID BIGINT REFERENCES posts(id) ON DELETE CASCADE,
    commentParentID BIGINT REFERENCES comments(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_comments_post_parent_id ON comments (postParentID) WHERE postParentID IS NOT NULL;
CREATE INDEX IF NOT EXISTS idx_comments_comment_parent_id ON comments (commentParentID) WHERE commentParentID IS NOT NULL;
