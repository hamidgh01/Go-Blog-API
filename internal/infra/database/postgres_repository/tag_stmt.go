package postgres_repository

import "database/sql"

const (
	// create
	bulkCreateTagQuery = `
		INSERT INTO tags (name) VALUES %s ON CONFLICT DO NOTHING
		RETURNING id, name
	`

	// read
	getTagByIDQuery   = "SELECT id, name FROM tags WHERE id = $1"
	getTagByNameQuery = "SELECT id, name FROM tags WHERE name = $1"

	getListOfTagsByNamesQuery = "SELECT id, name FROM tags WHERE name IN (%s)"
)

var (
	// read
	getTagByIDStmt   *sql.Stmt
	getTagByNameStmt *sql.Stmt
)

func prepareAllTagStatements(db *sql.DB) {
	getTagByIDStmt = prepareStatement(db, "getTagByID", getTagByIDQuery)
	getTagByNameStmt = prepareStatement(db, "getTagByName", getTagByNameQuery)
}

func closeAllTagStatements() {
	getTagByIDStmt.Close()
	getTagByNameStmt.Close()
}
