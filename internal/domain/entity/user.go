package entity

import (
	"database/sql"
	"time"
)

// Entity(Table): User (users)
// Relations:
// _ 1:N (One to Many) with 'Link'
// _ 1:N (One to Many) with 'Post'
// _ N:N (Many to Many) with 'Post' (like-system) -> via 'PostLikesM2M' association table
// _ 1:N (One to Many) with 'Comment'
// _ 1:N (One to Many) with 'List' (owned-lists)
// _ N:N (Many to Many) with 'List' (saved-lists) -> via 'UsersSavedListsM2M' association table
// _ N:N (Many to Many) with 'User' itself (follow-system) -> via 'FollowsM2M' association table
type User struct {
	ID          uint64       // sql: BIGSERIAL PRIMARY KEY (automatically indexed)
	Username    string       // sql: VARCHAR(64) NOT NULL UNIQUE (automatically indexed)
	Email       string       // sql: VARCHAR NOT NULL UNIQUE (automatically indexed)
	Password    string       // sql: VARCHAR NOT NULL
	Bio         string       // sql: VARCHAR(500)
	Enabled     bool         // sql: BOOLEAN DEFAULT true
	IsSuperuser bool         // sql: BOOLEAN DEFAULT false
	CreatedAt   time.Time    // sql: TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp
	ModifiedAt  sql.NullTime // sql: TIMESTAMP WITH TIME ZONE
}

// 'FollowsM2M' association table -> M2M between 'User' and 'User' (follow-system)
type FollowsM2M struct {
	FollowedBy uint64    // sql: BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE
	Followed   uint64    // sql: BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE (indexed -separately-)
	FollowedAt time.Time // sql: TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp
} // PRIMARY KEY (followed_by, followed)
