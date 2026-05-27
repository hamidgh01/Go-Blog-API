package entity

import (
	"database/sql"
	"time"
)

type PostStatus string

const (
	DR PostStatus = "draft"
	PB PostStatus = "published"
	RJ PostStatus = "rejected"
	DL PostStatus = "deleted-by-author"
)

// Entity(Table): Post (posts)
// Relations:
// _ N:1 (Many to One) with 'User'
// _ N:N (Many to Many) with 'User' (like-system) -> via 'PostLikesM2M' association table
// _ 1:N (One to Many) with 'Comment'
// _ N:N (Many to Many) with 'Tag' -> via 'PostTagsM2M' association table
// _ N:N (Many to Many) with 'List' -> via 'SavedPostsM2M' association table
type Post struct {
	ID               uint64       // sql: BIGSERIAL PRIMARY KEY (automatically indexed)
	Title            string       // sql: VARCHAR(200) NOT NULL (indexed)
	Content          string       // sql: TEXT
	Status           PostStatus   // sql: PostStatus DEFAULT 'draft' (PostStatus: created type as enum)
	IsPrivate        bool         // sql: BOOLEAN DEFAULT false
	CreatedAt        time.Time    // sql: TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp
	ModifiedAt       sql.NullTime // sql: TIMESTAMP WITH TIME ZONE
	FirstPublishedAt sql.NullTime // sql: TIMESTAMP WITH TIME ZONE
	UserID           uint64       // sql: BIGINT NOT NULL DEFAULT 0 REFERENCES users(id) ON DELETE SET DEFAULT
} // composite index: composite_idx_posts_userID_pubAt ON posts (userID, firstPublishedAt)

// 'PostLikesM2M' association table -> M2M between 'User' and 'Post' (like-system)
type PostLikesM2M struct {
	PostID  uint64    // sql: BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE
	UserID  uint64    // sql: BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE
	likedAt time.Time // sql: TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp
} // PRIMARY KEY (post_id, user_id)
