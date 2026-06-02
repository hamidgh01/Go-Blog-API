package entity

import (
	"database/sql"
	"time"
)

// Entity(Table): List (lists)
// Relations:
// _ N:1 (Many to One) with 'User' (owner)
// _ N:N (Many to Many) with 'User' (who-saved-this-list) -> via 'UsersSavedListsM2M' association table
// _ N:N (Many to Many) with 'Post' -> via 'SavedPostsM2M' association table
type List struct {
	ID          uint64       // sql: BIGSERIAL PRIMARY KEY (automatically indexed)
	Title       string       // sql: VARCHAR(100) NOT NULL
	Description string       // sql: VARCHAR(1000)
	IsPrivate   bool         // sql: BOOLEAN DEFAULT true
	CreatedAt   time.Time    // sql: TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp
	ModifiedAt  sql.NullTime // sql: TIMESTAMP WITH TIME ZONE
	UserID      uint64       // sql: BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE (indexed)
	// NOTE: `UserID` here is the `id` of the user who owns this list
	// FK:
	User *User
}

// 'SavedPostsM2M' association table -> M2M between 'List' and 'Post'
type SavedPostsM2M struct {
	ListID  uint64    // sql: BIGINT NOT NULL REFERENCES lists(id) ON DELETE CASCADE
	PostID  uint64    // sql: BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE (indexed -separately-)
	SavedAt time.Time // sql: TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp
} // PRIMARY KEY (list_id, post_id)

// 'UsersSavedListsM2M' association table -> M2M between 'User' and 'List'
type UsersSavedListsM2M struct {
	// NOTE: `UserID` here is the `id` of the user who saved a list
	UserID  uint64    // sql: BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE
	ListID  uint64    // sql: BIGINT NOT NULL REFERENCES lists(id) ON DELETE CASCADE
	SavedAt time.Time // sql: TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp
} // PRIMARY KEY (user_id, list_id)
