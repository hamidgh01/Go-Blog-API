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
