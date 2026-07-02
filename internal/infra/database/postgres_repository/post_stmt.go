package postgres_repository

import "database/sql"

var (
	// create
	createDraftPostQuery = `
		INSERT INTO posts (title, content, isPrivate, userID) VALUES ($1, $2, $3, $4)
		RETURNING id, title, content, status, isPrivate, userID, createdAt, modifiedAt, firstPublishedAt
	`

	createPublishedPostQuery = `
		INSERT INTO posts (title, content, status, isPrivate, firstPublishedAt, userID)
		VALUES ($1, $2, 'published', $3, CURRENT_TIMESTAMP, $4)
		RETURNING id, title, content, status, isPrivate, userID, createdAt, modifiedAt, firstPublishedAt
	`

	// update
	updatePostQuery = `
		UPDATE posts
		SET title = $1, content = $2, modifiedAt = CURRENT_TIMESTAMP
		WHERE id = $3 AND (status = 'draft' OR status = 'published')
	`
	updatePostStatusQuery = `
		UPDATE posts
		SET status = $1, modifiedAt = CURRENT_TIMESTAMP
		WHERE id = $2
	`
	publishDraftPostQuery = `
		UPDATE posts
		SET status = 'published', modifiedAt = CURRENT_TIMESTAMP, firstPublishedAt = CURRENT_TIMESTAMP
		WHERE id = $1 AND status = 'draft'
	`
	updatePostPrivacyQuery = `
		UPDATE posts
		SET isPrivate = $1, modifiedAt = CURRENT_TIMESTAMP
		WHERE id = $2 AND (status = 'draft' OR status = 'published')
	`

	// delete
	deletePostQuery = "DELETE FROM posts WHERE id = $1"

	// read
	getPostByIDQuery = `
		SELECT
		p.id, p.title, p.content, p.status, p.isPrivate, p.userID, p.createdAt, p.modifiedAt, p.firstPublishedAt,
		u.id, u.username
		FROM posts as p
		JOIN users as u
		ON p.userID = u.id
		WHERE p.id = $1 AND status = 'published'
	`

	getDraftByIDQuery = `` // add later (and implement other layers) access: admin_or_owner

	getPostOwnerIDQuery = "SELECT userID FROM posts WHERE id = $1"
)

var (
	// create
	createDraftPostStmt     *sql.Stmt
	createPublishedPostStmt *sql.Stmt

	// update
	updatePostStmt        *sql.Stmt
	updatePostStatusStmt  *sql.Stmt
	publishDraftPostStmt  *sql.Stmt
	updatePostPrivacyStmt *sql.Stmt

	// delete
	deletePostStmt *sql.Stmt

	// read
	getPostByIDStmt *sql.Stmt

	getPostOwnerIDStmt *sql.Stmt
)

func prepareAllPostStatements(db *sql.DB) {
	// create
	createDraftPostStmt = prepareStatement(db, "createDraftPost", createDraftPostQuery)
	createPublishedPostStmt = prepareStatement(db, "createPublishedPost", createPublishedPostQuery)

	// update
	updatePostStmt = prepareStatement(db, "updatePost", updatePostQuery)
	updatePostStatusStmt = prepareStatement(db, "updatePostStatus", updatePostStatusQuery)
	publishDraftPostStmt = prepareStatement(db, "publishDraftPost", publishDraftPostQuery)
	updatePostPrivacyStmt = prepareStatement(db, "updatePostPrivacy", updatePostPrivacyQuery)

	// delete
	deletePostStmt = prepareStatement(db, "deletePost", deletePostQuery)

	// read
	getPostByIDStmt = prepareStatement(db, "getPostByID", getPostByIDQuery)

	getPostOwnerIDStmt = prepareStatement(db, "getPostOwnerID", getPostOwnerIDQuery)
}

func closeAllPostStatements() {
	createDraftPostStmt.Close()
	createPublishedPostStmt.Close()
	updatePostStmt.Close()
	updatePostStatusStmt.Close()
	publishDraftPostStmt.Close()
	updatePostPrivacyStmt.Close()
	deletePostStmt.Close()
	getPostByIDStmt.Close()
	getPostOwnerIDStmt.Close()
}
