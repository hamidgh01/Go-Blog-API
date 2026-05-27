package entity

import (
	"database/sql"
	"time"
)

type List struct {
	ID          uint64       // sql: BIGSERIAL PRIMARY KEY (automatically indexed)
	Title       string       // sql: VARCHAR(100) NOT NULL
	Description string       // sql: VARCHAR(1000)
	IsPrivate   bool         // sql: BOOLEAN DEFAULT true
	CreatedAt   time.Time    // sql: TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp
	ModifiedAt  sql.NullTime // sql: TIMESTAMP WITH TIME ZONE
	UserID      uint64       // sql: BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE (indexed)
}

type SavedPostsM2M struct {
	ListID  uint64    // sql: BIGINT NOT NULL REFERENCES lists(id) ON DELETE CASCADE
	PostID  uint64    // sql: BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE (indexed -separately-)
	SavedAt time.Time // sql: TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp
} // PRIMARY KEY (list_id, post_id)

type UsersSavedListsM2M struct {
	UserID  uint64    // sql: BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE
	ListID  uint64    // sql: BIGINT NOT NULL REFERENCES lists(id) ON DELETE CASCADE
	SavedAt time.Time // sql: TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp
} // PRIMARY KEY (user_id, list_id)
