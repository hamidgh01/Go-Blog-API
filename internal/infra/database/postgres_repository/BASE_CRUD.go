package postgres_repository

import (
	"context"
	"database/sql"
	"fmt"

	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	dbErrors "github.com/hamidgh01/Go-Blog-API/internal/infra/database/errors"
)

func createOrDeleteM2MRelationship(
	ctx context.Context,
	createOrDeleteM2MRelationshipStmt *sql.Stmt,
	firstID uint64,
	secondID uint64,
	noRowsAffectedMessage string,
) error {
	result, err := createOrDeleteM2MRelationshipStmt.ExecContext(ctx, firstID, secondID)
	if err != nil {
		return dbErrors.GetDBError(err)
	}

	n, e := result.RowsAffected()
	if n == 0 || e == sql.ErrNoRows {
		return dbErrors.NewNoRowsAffectedOnM2MEntity(noRowsAffectedMessage)
	} else if e != nil {
		return dbErrors.GetDBError(e)
	}

	return nil
}

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

func getListOfOuterResourceByFK(
	ctx context.Context,
	db *sql.DB,
	fk uint64,
	page *d.PaginationQueryParams,
	countQuery string,
	mainQuery string,
	operationName string,
	notFoundMessage string,
) (rows *sql.Rows, totalRows, pageNum, pageSize, totalPages int, err error) {

	err = db.QueryRowContext(ctx, countQuery, fk).Scan(&totalRows)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, 0, 0, 0, 0, dbErrors.NewRecordNotFoundError(
				fmt.Sprintf("%s with id='%d'", notFoundMessage, fk),
			)
		}
		return nil, 0, 0, 0, 0, dbErrors.GetDBError(err)
	}

	if totalRows == 0 {
		return nil, 0, 0, 0, 0, dbErrors.NewRecordNotFoundError(
			fmt.Sprintf("%s with id='%d'", notFoundMessage, fk),
		)
	}

	totalPages = max((totalRows+page.GetSize()-1)/page.GetSize(), 1)
	// e.g. totalRows=30, Size = 10 ->  max(39 / 10, 1) ->  max (3, 1) ->  TotalPages=3
	// e.g. totalRows=31, Size = 10 ->  max(40 / 10, 1) ->  max (4, 1) ->  TotalPages=4
	// e.g. totalRows=8, Size = 10  ->  max(17 / 10, 1) ->  max (1, 1) ->  TotalPages=1

	// if page-number is greater than total pages -> set it to last page
	if page.GetPage() > totalPages {
		page.Page = totalPages
	}

	query := fmt.Sprintf(mainQuery, page.GetSize(), page.GetOffset())

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return nil, 0, 0, 0, 0, fmt.Errorf(
			"failed to prepare %s stmt. origin: %w", operationName, err,
		)
	}
	defer stmt.Close()

	rows, err = stmt.QueryContext(ctx, fk)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, 0, 0, 0, 0, dbErrors.NewRecordNotFoundError(
				fmt.Sprintf("%s (id='%d')", notFoundMessage, fk),
			)
		}
		return nil, 0, 0, 0, 0, dbErrors.GetDBError(err)
	}

	return rows, totalRows, page.GetPage(), page.GetSize(), totalPages, nil
}
