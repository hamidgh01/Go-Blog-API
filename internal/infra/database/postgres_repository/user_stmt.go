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
