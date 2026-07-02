package postgres_repository

import "database/sql"

const (
	// create
	createListQuery = `
		INSERT INTO lists (title, description, isPrivate, userID) VALUES ($1, $2, $3, $4)
		RETURNING id, title, description, isPrivate, userID, createdAt, modifiedAt
	`

	// update
	updateListQuery = `
		UPDATE lists
		SET title = $1, description = $2, isPrivate = $3, modifiedAt = CURRENT_TIMESTAMP
		WHERE id = $4
	`
	updateListPrivacyQuery = `
		UPDATE lists
		SET isPrivate = $1, modifiedAt = CURRENT_TIMESTAMP
		WHERE id = $2
	`

	// delete
	deleteListQuery = "DELETE FROM lists WHERE id = $1"

	// read
	getListByIDQuery = `
		SELECT
		l.id, l.title, l.description, l.isPrivate, l.userID, l.createdAt, l.modifiedAt,
		u.id, u.username
		FROM lists as l
		JOIN users as u
		ON l.userID = u.id
		WHERE l.id = $1
	`

	getListOwnerIDQuery = "SELECT userID FROM lists WHERE id = $1"

	// get other sources that has FK to `List`
	countSavedPostsQuery    = "SELECT COUNT(post_id) FROM saved_posts_m2m WHERE list_id = $1"
	countUsersWhoSavedQuery = "SELECT COUNT(user_id) FROM users_saved_lists_m2m WHERE list_id = $1"

	getSavedPostsQuery = `
		SELECT
		p.id, p.title, p.isPrivate, p.userID, p.createdAt, p.modifiedAt, p.firstPublishedAt,
		u.id, u.username
		FROM posts as p
		JOIN users as u ON p.userID = u.id
		WHERE p.id IN (
			SELECT sp.post_id FROM saved_posts_m2m as sp
			WHERE sp.list_id = $1
			ORDER BY sp.saved_at DESC
			LIMIT %d OFFSET %d
		) AND p.status = 'published' AND p.isPrivate = false
	`
	getUsersWhoSavedQuery = `
		SELECT u.id, u.username
		FROM users as u
		WHERE u.id IN (
			SELECT usl.user_id FROM users_saved_lists_m2m as usl
			WHERE usl.list_id = $1
			ORDER BY usl.saved_at DESC
			LIMIT %d OFFSET %d
		) AND enabled = true
	`
)

var (
	// create
	createListStmt *sql.Stmt

	// update
	updateListStmt        *sql.Stmt
	updateListPrivacyStmt *sql.Stmt

	// delete
	deleteListStmt *sql.Stmt

	// read
	getListByIDStmt *sql.Stmt

	getListOwnerIDStmt *sql.Stmt
)

func prepareAllListStatements(db *sql.DB) {
	// create
	createListStmt = prepareStatement(db, "createList", createListQuery)

	// update
	updateListStmt = prepareStatement(db, "updateList", updateListQuery)
	updateListPrivacyStmt = prepareStatement(db, "updateListPrivacy", updateListPrivacyQuery)

	// delete
	deleteListStmt = prepareStatement(db, "deleteList", deleteListQuery)

	// read
	getListByIDStmt = prepareStatement(db, "getListByID", getListByIDQuery)

	getListOwnerIDStmt = prepareStatement(db, "getListOwnerID", getListOwnerIDQuery)
}

func closeAllListStatements() {
	createListStmt.Close()
	updateListStmt.Close()
	updateListPrivacyStmt.Close()
	deleteListStmt.Close()
	getListByIDStmt.Close()
	getListOwnerIDStmt.Close()
}
