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

	// get other sources that has FK to `Tag` (just associated posts)
	countAssociatedPostsQuery = "SELECT COUNT(post_id) FROM posts_tags_m2m WHERE tag_id = $1"
	getAssociatedPostsQuery   = `
		SELECT
		p.id, p.title, p.isPrivate, p.userID, p.createdAt, p.modifiedAt, p.firstPublishedAt,
		u.id, u.username
		FROM posts as p
		JOIN users as u ON p.userID = u.id
		WHERE p.id IN (
			SELECT pt.post_id FROM posts_tags_m2m as pt
			WHERE pt.tag_id = $1
			LIMIT %d OFFSET %d
		) AND p.status = 'published' AND p.isPrivate = false
	`
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
