package entity

import (
	"database/sql"
	"time"
)

type CommentStatus string

const (
	PUB CommentStatus = "published"
	HID CommentStatus = "hidden-by-Admin"
	DEL CommentStatus = "deleted-by-commenter"
)

// Entity(Table): Comment (comments)
// Relations:
// _ N:1 (Many to One) with 'User'
// _ N:1 (Many to One) with 'Post'
// _ Self-Referenced with 'Comment' itself (a `comment` to a post can have some `replies`)
// Notes:
// parent of a comment can be another 'comment' or a 'post'. so for a comment:
// if parent is a 'post' --> CommentParentID will be NULL
// if parent is another 'comment' --> PostParentID will be NULL
// (this is why CommentParentID & PostParentID are nullable)
type Comment struct {
	ID              uint64       // sql: BIGSERIAL PRIMARY KEY (automatically indexed)
	Content         string       // sql: VARCHAR(1500) NOT NULL
	Status          PostStatus   // sql: CommentStatus DEFAULT 'published'
	CreatedAt       time.Time    // sql: TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp
	ModifiedAt      sql.NullTime // sql: TIMESTAMP WITH TIME ZONE
	UserID          uint64       // sql: BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE
	PostParentID    uint64       // sql: BIGINT REFERENCES posts(id) ON DELETE CASCADE (indexed WHERE postParentID IS NOT NULL;)
	CommentParentID uint64       // sql: BIGINT REFERENCES comments(id) ON DELETE CASCADE (indexed WHERE commentParentID IS NOT NULL;)
}
