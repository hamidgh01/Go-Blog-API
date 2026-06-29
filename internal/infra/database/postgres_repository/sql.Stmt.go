package postgres_repository

import (
	"database/sql"

	"github.com/hamidgh01/Go-Blog-API/pkg/logging"
)

func prepareStatement(db *sql.DB, stmtName string, query string) *sql.Stmt {
	logger := logging.GetLogger()

	stmt, err := db.Prepare(query)
	if err != nil {
		logger.Fatalf("failed to prepare %s stmt. reason: %s", stmtName, err.Error())
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
	// M2M entities
	closeAllFollowStatements()
	closeAllLikeStatements()
	closeAllSavePostStatements()
	closeAllSaveListStatements()
}
