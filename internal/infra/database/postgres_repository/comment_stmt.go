package postgres_repository

import (
	"database/sql"
	"fmt"
)

const DETAILED_COMMENT_FIELDS string = "id, content, status, postParentID, commentParentID, userID, createdAt, modifiedAt"

var (
	// create
	createCommentQuery string = fmt.Sprintf(
		`INSERT INTO comments (content, postParentID, userID) VALUES ($1, $2, $3)
		RETURNING %s`, DETAILED_COMMENT_FIELDS,
	)

	createReplyQuery string = fmt.Sprintf(
		`INSERT INTO comments (content, commentParentID, userID) VALUES ($1, $2, $3)
		RETURNING %s`, DETAILED_COMMENT_FIELDS,
	)

	// update
	updateCommentQuery = `
		UPDATE comments
		SET content = $1, modifiedAt = CURRENT_TIMESTAMP
		WHERE id = $2 AND status = 'published'
	`
	updateCommentStatusQuery = `
		UPDATE comments
		SET status = $1, modifiedAt = CURRENT_TIMESTAMP
		WHERE id = $2
	`

	// delete
	deleteCommentQuery = "DELETE FROM comments WHERE id = $1"

	// read
	getCommentByIDQuery = `
		SELECT
		c.id, c.content, c.status, c.postParentID, c.commentParentID, c.userID, c.createdAt, c.modifiedAt,
		u.id, u.username
		FROM comments as c
		JOIN users as u
		ON c.userID = u.id
		WHERE c.id = $1
	`

	getCommentOwnerIDQuery = "SELECT userID FROM comments WHERE id = $1"
)

var (
	// create
	createCommentStmt *sql.Stmt
	createReplyStmt   *sql.Stmt

	// update
	updateCommentStmt       *sql.Stmt
	updateCommentStatusStmt *sql.Stmt

	// delete
	deleteCommentStmt *sql.Stmt

	// read
	getCommentByIDStmt *sql.Stmt

	getCommentOwnerIDStmt *sql.Stmt
)

func prepareAllCommentStatements(db *sql.DB) {
	// create
	createCommentStmt = prepareStatement(db, "createComment", createCommentQuery)
	createReplyStmt = prepareStatement(db, "createReply", createReplyQuery)

	// update
	updateCommentStmt = prepareStatement(db, "updateComment", updateCommentQuery)
	updateCommentStatusStmt = prepareStatement(db, "updateCommentStatus", updateCommentStatusQuery)

	// delete
	deleteCommentStmt = prepareStatement(db, "deleteComment", deleteCommentQuery)

	// read
	getCommentByIDStmt = prepareStatement(db, "getCommentByID", getCommentByIDQuery)

	getCommentOwnerIDStmt = prepareStatement(db, "getCommentOwnerID", getCommentOwnerIDQuery)
}

func closeAllCommentStatements() {
	createCommentStmt.Close()
	createReplyStmt.Close()
	updateCommentStmt.Close()
	updateCommentStatusStmt.Close()
	deleteCommentStmt.Close()
	getCommentByIDStmt.Close()
	getCommentOwnerIDStmt.Close()
}
