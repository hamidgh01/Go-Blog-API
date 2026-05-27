package entity

// Entity(Table): Tag (tags)
// Relations:
// _ N:N (Many to Many) with 'Post' -> via 'PostTagsM2M' association table
type Tag struct {
	ID   uint64 // sql: BIGSERIAL PRIMARY KEY (automatically indexed)
	Name string // sql: VARCHAR(32) NOT NULL UNIQUE (automatically indexed)
}

// 'PostTagsM2M' association table -> M2M between 'Post' and 'Tag'
type PostTagsM2M struct {
	TagID  uint64 // sql: BIGINT NOT NULL REFERENCES tags(id) ON DELETE CASCADE
	PostID uint64 // sql: BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE (indexed -separately-)
} // PRIMARY KEY (tagID, postID)
