package entity

import (
	"database/sql"
	"time"
)

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

type FollowsM2M struct {
	FollowedBy uint64    // sql: BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE
	Followed   uint64    // sql: BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE (indexed -separately-)
	FollowedAt time.Time // sql: TIMESTAMP WITH TIME ZONE DEFAULT current_timestamp
} // PRIMARY KEY (followed_by, followed)
