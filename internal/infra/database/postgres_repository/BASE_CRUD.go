package postgres_repository

import (
	"context"
	"database/sql"
	"fmt"

	dbErrors "github.com/hamidgh01/Go-Blog-API/internal/infra/database/errors"
)

func update(ctx context.Context, updateStmt *sql.Stmt, entityName string, pk uint64, fields ...any) error {
	fields = append(fields, pk)
	result, err := updateStmt.ExecContext(ctx, fields...)
	if err != nil {
		if err == sql.ErrNoRows {
			return dbErrors.NewRecordNotFoundError(fmt.Sprintf("%s (with id=%d) not found", entityName, pk))
		}
		return dbErrors.GetDBError(err)
	}

	n, e := result.RowsAffected()
	if n == 0 || e == sql.ErrNoRows {
		return dbErrors.NewRecordNotFoundError(fmt.Sprintf("%s (with id=%d) not found", entityName, pk))
	} else if e != nil {
		return dbErrors.GetDBError(e)
	}

	return nil
}

func delete(ctx context.Context, deleteStmt *sql.Stmt, entityName string, pk uint64) error {
	result, err := deleteStmt.ExecContext(ctx, pk)
	if err != nil {
		if err == sql.ErrNoRows {
			return dbErrors.NewRecordNotFoundError(fmt.Sprintf("%s (with id=%d) not found", entityName, pk))
		}
		return dbErrors.GetDBError(err)
	}

	n, e := result.RowsAffected()
	if n == 0 || e == sql.ErrNoRows {
		return dbErrors.NewRecordNotFoundError(fmt.Sprintf("%s (with id=%d) not found", entityName, pk))
	} else if e != nil {
		return dbErrors.GetDBError(e)
	}

	return nil
}

func checkUniqueFieldExists(
	ctx context.Context, checkExistsStmt *sql.Stmt, uniqueFieldVal string,
) (bool, error) {
	var id uint64
	err := checkExistsStmt.QueryRowContext(ctx, uniqueFieldVal).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, dbErrors.GetDBError(err)
	}

	if id == 0 {
		return false, nil
	}

	return true, nil
}

func getFieldValue(
	ctx context.Context, getFieldValueStmt *sql.Stmt, entityName string, pk uint64,
) (any, error) {
	var result any
	err := getFieldValueStmt.QueryRowContext(ctx, pk).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, dbErrors.NewRecordNotFoundError(fmt.Sprintf("%s (with id=%d) not found", entityName, pk))
		}
		return false, dbErrors.GetDBError(err)
	}

	return result, nil
}

func getOwnerID(
	ctx context.Context, getOwnerIDStmt *sql.Stmt, entityName string, pk uint64,
) (ownerID uint64, err error) {
	err = getOwnerIDStmt.QueryRowContext(ctx, pk).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, dbErrors.NewRecordNotFoundError(fmt.Sprintf("%s (with id=%d) not found", entityName, pk))
		}
		return 0, dbErrors.GetDBError(err)
	}

	// if ownerID == 0 {
	// 	in this situation: related `user` (owner) is deleted and `resource.userId` is set to 0
	// 	handle it later
	// }

	return
}
