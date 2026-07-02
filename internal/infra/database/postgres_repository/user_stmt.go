package postgres_repository

import (
	"database/sql"
	"fmt"
)

const (
	BRIEF_USER_FIELDS    string = "id, username"
	DETAILED_USER_FIELDS string = "id, username, email, bio, enabled, createdAt, modifiedAt"
)

var (
	// create
	createUserQuery = fmt.Sprintf(
		`INSERT INTO users (username, email, password) VALUES ($1, $2, $3)
		RETURNING %s`, DETAILED_USER_FIELDS,
	)

	// update
	updateSingleUserFieldQuery = `
		UPDATE users
		SET %s = $1, modifiedAt = CURRENT_TIMESTAMP
		WHERE id = $2
	`
	updateUsernameQuery = fmt.Sprintf(updateSingleUserFieldQuery, "username")
	updateEmailQuery    = fmt.Sprintf(updateSingleUserFieldQuery, "email")
	updateBioQuery      = fmt.Sprintf(updateSingleUserFieldQuery, "bio")
	updatePasswordQuery = fmt.Sprintf(updateSingleUserFieldQuery, "password")
	updateEnabledQuery  = fmt.Sprintf(updateSingleUserFieldQuery, "enabled")

	// delete
	deleteUserQuery = "DELETE FROM users WHERE id = $1"

	// checking and others
	checkUsernameExistsQuery     = "SELECT id FROM users WHERE username = $1"
	checkEmailExistsQuery        = "SELECT id FROM users WHERE email = $1"
	checkIsEnabledQuery          = "SELECT enabled FROM users WHERE id = $1"
	checkIsSuperuserQuery        = "SELECT isSuperuser FROM users WHERE id = $1"
	getHashedPasswordQuery       = "SELECT password FROM users WHERE id = $1"
	getUserForLoginVerificationQ = "SELECT id, username, enabled, password FROM users WHERE email = $1 OR username = $1"

	// read
	getUserByIDQuery       = fmt.Sprintf("SELECT %s FROM users WHERE id = $1", DETAILED_USER_FIELDS)
	getUserByUsernameQuery = fmt.Sprintf("SELECT %s FROM users WHERE username = $1", DETAILED_USER_FIELDS)
	getUserByEmailQuery    = fmt.Sprintf("SELECT %s FROM users WHERE email = $1", DETAILED_USER_FIELDS)

	// get other sources that has FK to `User`

	// count queries
	countFollowersQuery    = "SELECT COUNT(followed_by) FROM follows_m2m WHERE followed = $1"
	countFollowingsQuery   = "SELECT COUNT(followed) FROM follows_m2m WHERE followed_by = $1"
	countUserPostsQuery    = "SELECT COUNT(id) FROM posts WHERE userID = $1 AND status = 'published' AND isPrivate = false"
	countUserLikesQuery    = "SELECT COUNT(post_id) FROM post_likes_m2m WHERE user_id = $1"
	countUserCommentsQuery = "SELECT COUNT(id) FROM comments WHERE userID = $1 AND status = 'published'"
	countUserLinksQuery    = "SELECT COUNT(id) FROM links WHERE userID = $1"
	countOwnedListsQuery   = "SELECT COUNT(id) FROM lists WHERE userID = $1 AND isPrivate = false"
	countSavedListsQuery   = "SELECT COUNT(list_id) FROM users_saved_lists_m2m WHERE user_id = $1"

	getFollowersByUserIdFkQuery = `
		SELECT u.id, u.username
		FROM users as u
		WHERE u.id IN (
			SELECT f.followed_by FROM follows_m2m as f
			WHERE f.followed = $1
			ORDER BY f.followed_at DESC
			LIMIT %d OFFSET %d
		)
	`
	getFollowingsByUserIdFkQuery = `
		SELECT u.id, u.username
		FROM users as u
		WHERE u.id IN (
			SELECT f.followed FROM follows_m2m as f
			WHERE f.followed_by = $1
			ORDER BY f.followed_at DESC
			LIMIT %d OFFSET %d
		)
	`
	getPostsByUserIdFkQuery = `
		SELECT
		p.id, p.title, p.isPrivate, p.userID, p.createdAt, p.modifiedAt, p.firstPublishedAt,
		u.id, u.username
		FROM posts as p
		JOIN users as u ON p.userID = u.id
		WHERE p.userID = $1 AND p.status = 'published' AND p.isPrivate = false
		ORDER BY firstPublishedAt DESC
		LIMIT %d OFFSET %d
	`
	getLikesByUserIdFkQuery = `
		SELECT
		p.id, p.title, p.isPrivate, p.userID, p.createdAt, p.modifiedAt, p.firstPublishedAt,
		u.id, u.username
		FROM posts as p
		JOIN users as u ON p.userID = u.id
		WHERE p.id IN (
			SELECT pl.post_id FROM post_likes_m2m as pl
			WHERE pl.user_id = $1
			ORDER BY pl.liked_at DESC
			LIMIT %d OFFSET %d
		) AND p.status = 'published' AND p.isPrivate = false
	`
	getCommentsByUserIdFkQuery = `
		SELECT
		c.id, c.content, c.status, c.postParentID, c.commentParentID, c.userID, c.createdAt, c.modifiedAt,
		u.id, u.username
		FROM comments as c
		JOIN users as u ON c.userID = u.id
		WHERE c.userID = $1 AND c.status = 'published'
		ORDER BY c.createdAt DESC
		LIMIT %d OFFSET %d
	`
	getLinksByUserIdFkQuery = `
		SELECT l.id, l.title, l.url, l.userID, u.id, u.username
		FROM links as l
		JOIN users as u ON l.userID = u.id
		WHERE l.userID = $1
		ORDER BY l.id DESC
		LIMIT %d OFFSET %d
	`
	getOwnedListsByUserIdFkQuery = `
		SELECT
		l.id, l.title, l.isPrivate, l.userID, l.createdAt, l.modifiedAt,
		u.id, u.username
		FROM lists as l
		JOIN users as u ON l.userID = u.id
		WHERE l.userID = $1 AND l.isPrivate = false
		ORDER BY l.createdAt DESC
		LIMIT %d OFFSET %d
	`
	getSavedListsByUserIdFkQuery = `
		SELECT
		l.id, l.title, l.isPrivate, l.userID, l.createdAt, l.modifiedAt,
		u.id, u.username
		FROM lists as l
		JOIN users as u ON l.userID = u.id
		WHERE l.id IN (
			SELECT usl.list_id FROM users_saved_lists_m2m as usl
			WHERE usl.user_id = $1
			ORDER BY usl.saved_at DESC
			LIMIT %d OFFSET %d
		) AND l.isPrivate = false
	`
)

var (
	// create
	createUserStmt *sql.Stmt

	// update
	updateUsernameStmt *sql.Stmt
	updateEmailStmt    *sql.Stmt
	updateBioStmt      *sql.Stmt
	updatePasswordStmt *sql.Stmt
	updateEnabledStmt  *sql.Stmt

	// delete
	deleteUserStmt *sql.Stmt

	// checking
	checkUsernameExistsStmt         *sql.Stmt
	checkEmailExistsStmt            *sql.Stmt
	checkIsEnabledStmt              *sql.Stmt
	checkIsSuperuserStmt            *sql.Stmt
	getHashedPasswordStmt           *sql.Stmt
	getUserForLoginVerificationStmt *sql.Stmt

	// read
	getUserByIDStmt       *sql.Stmt
	getUserByUsernameStmt *sql.Stmt
	getUserByEmailStmt    *sql.Stmt
)

func prepareAllUserStatements(db *sql.DB) {
	// create
	createUserStmt = prepareStatement(db, "createUser", createUserQuery)

	// update
	updateUsernameStmt = prepareStatement(db, "updateUserUsername", updateUsernameQuery)
	updateEmailStmt = prepareStatement(db, "updateUserEmail", updateEmailQuery)
	updateBioStmt = prepareStatement(db, "updateUserBio", updateBioQuery)
	updatePasswordStmt = prepareStatement(db, "updateUserPassword", updatePasswordQuery)
	updateEnabledStmt = prepareStatement(db, "updateUserEnabled", updateEnabledQuery)

	// delete
	deleteUserStmt = prepareStatement(db, "deleteUser", deleteUserQuery)

	// checking
	checkUsernameExistsStmt = prepareStatement(db, "checkUsernameExists", checkUsernameExistsQuery)
	checkEmailExistsStmt = prepareStatement(db, "checkEmailExists", checkEmailExistsQuery)
	checkIsEnabledStmt = prepareStatement(db, "checkIsEnabled", checkIsEnabledQuery)
	checkIsSuperuserStmt = prepareStatement(db, "checkIsSuperuser", checkIsSuperuserQuery)
	getHashedPasswordStmt = prepareStatement(db, "getHashedPassword", getHashedPasswordQuery)
	getUserForLoginVerificationStmt = prepareStatement(db, "getUserForLoginVerification", getUserForLoginVerificationQ)

	// read
	getUserByIDStmt = prepareStatement(db, "getUserByID", getUserByIDQuery)
	getUserByUsernameStmt = prepareStatement(db, "getUserByUsername", getUserByUsernameQuery)
	getUserByEmailStmt = prepareStatement(db, "getUserByEmail", getUserByEmailQuery)
}

func closeAllUserStatements() {
	createUserStmt.Close()
	updateUsernameStmt.Close()
	updateEmailStmt.Close()
	updateBioStmt.Close()
	updatePasswordStmt.Close()
	updateEnabledStmt.Close()
	deleteUserStmt.Close()
	checkUsernameExistsStmt.Close()
	checkEmailExistsStmt.Close()
	checkIsEnabledStmt.Close()
	checkIsSuperuserStmt.Close()
	getHashedPasswordStmt.Close()
	getUserForLoginVerificationStmt.Close()
	getUserByIDStmt.Close()
	getUserByUsernameStmt.Close()
	getUserByEmailStmt.Close()
}
