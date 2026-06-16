package postgres_repository

import (
	"database/sql"
	"fmt"
	"os"
)

func prepareStatement(db *sql.DB, stmtName string, query string) *sql.Stmt {
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Printf("failed to prepare %s stmt. reason: %s", stmtName, err.Error())
		os.Exit(1)
	}
	return stmt
}

func CloseAllPreparedStatements() {
	closeAllUserStatements()
	closeAllPostStatements()
	closeAllCommentStatements()
	closeAllListStatements()
	closeAllLinkStatements()
	closeAllTagStatements()
}
