package postgres_repository

import "database/sql"

const (
	// create
	createLinkQuery = `
		INSERT INTO links (title, url, userID) VALUES ($1, $2, $3)
		RETURNING id, title, url, userID
	`

	// update
	updateLinkQuery = `
		UPDATE links
		SET title = $1, url = $2
		WHERE id = $3
	`

	// delete
	deleteLinkQuery = "DELETE FROM links WHERE id = $1"

	// read
	getLinkByIDQuery = `
		SELECT l.id, l.title, l.url, l.userID, u.id, u.username
		FROM links as l
		JOIN users as u
		ON l.userID = u.id
		WHERE l.id = $1
	`

	getLinkOwnerIDQuery = "SELECT userID FROM links WHERE id = $1"
)

var (
	// create
	createLinkStmt *sql.Stmt

	// update
	updateLinkStmt *sql.Stmt

	// delete
	deleteLinkStmt *sql.Stmt

	// read
	getLinkByIDStmt *sql.Stmt

	getLinkOwnerIDStmt *sql.Stmt
)

func prepareAllLinkStatements(db *sql.DB) {
	// create
	createLinkStmt = prepareStatement(db, "createLink", createLinkQuery)

	// update
	updateLinkStmt = prepareStatement(db, "updateLink", updateLinkQuery)

	// delete
	deleteLinkStmt = prepareStatement(db, "deleteLink", deleteLinkQuery)

	// read
	getLinkByIDStmt = prepareStatement(db, "getLinkByID", getLinkByIDQuery)

	getLinkOwnerIDStmt = prepareStatement(db, "getLinkOwnerID", getLinkOwnerIDQuery)
}

func closeAllLinkStatements() {
	createLinkStmt.Close()
	updateLinkStmt.Close()
	deleteLinkStmt.Close()
	getLinkByIDStmt.Close()
	getLinkOwnerIDStmt.Close()
}
