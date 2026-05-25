package entity

type Tag struct {
	ID   uint64 // sql: BIGSERIAL PRIMARY KEY (automatically indexed)
	Name string // sql: VARCHAR(32) NOT NULL UNIQUE (automatically indexed)
}

type PostTagsM2M struct {
	TagID  uint64 // sql: BIGINT NOT NULL REFERENCES tags(id) ON DELETE CASCADE
	PostID uint64 // sql: BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE (indexed -separately-)
} // PRIMARY KEY (tagID, postID)
