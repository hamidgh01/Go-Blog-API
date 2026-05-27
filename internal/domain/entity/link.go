package entity

// Entity(Table): Link (links)
// Relations:
// _ N:1 (Many to One) with 'User'
type Link struct {
	ID     uint64 // sql: BIGSERIAL PRIMARY KEY (automatically indexed)
	Title  string // sql: VARCHAR(32) NOT NULL
	Url    string // sql: VARCHAR NOT NULL
	UserID uint64 // sql: BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE (indexed)
}
