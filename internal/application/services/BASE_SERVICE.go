package services

import (
	"context"
	"fmt"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/http/generics"
)

func createOrDeleteM2MRelationship[TEntity e.TM2MDBEntities](
	ctx context.Context,
	serviceName string,
	entity *TEntity,
	createOrDeleteM2MRelationshipRepo func(ctx context.Context, entity *TEntity) error,
) *service_errors.ServiceError {
	if err := createOrDeleteM2MRelationshipRepo(ctx, entity); err != nil {
		return service_errors.MapDBErrToServiceErr(err, serviceName)
	}

	return nil
}

func create[TEntity e.TDBEntities, TOutput generics.OutputTypes](
	ctx context.Context,
	objectName string,
	entity *TEntity,
	createRepository func(ctx context.Context, entity *TEntity) (*TEntity, error),
	toObjectDetailsDTOFunc func(entity *TEntity) *TOutput,
) (*TOutput, *service_errors.ServiceError) {
	//
	createdObj, err := createRepository(ctx, entity)
	if err != nil {
		return nil, service_errors.MapDBErrToServiceErr(err, fmt.Sprintf("create %s", objectName))
	}

	//
	return toObjectDetailsDTOFunc(createdObj), nil
}

func update[TEntity e.TDBEntities](
	ctx context.Context,
	pk uint64,
	objectName string,
	entity *TEntity,
	updateRepository func(ctx context.Context, pk uint64, entity *TEntity) error,
) *service_errors.ServiceError {
	err := updateRepository(ctx, pk, entity)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, fmt.Sprintf("update %s", objectName))
	}

	return nil
}

func delete(
	ctx context.Context,
	pk uint64,
	objectName string,
	deleteRepository func(ctx context.Context, pk uint64) error,
) *service_errors.ServiceError {
	err := deleteRepository(ctx, pk)
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, fmt.Sprintf("delete %s", objectName))
	}

	return nil
}

func getByID[TEntity e.TDBEntities, TOutput generics.OutputTypes](
	ctx context.Context,
	pk uint64,
	objectName string,
	getByIdRepository func(ctx context.Context, pk uint64) (*TEntity, error),
	toObjectDetailsDTOFunc func(entity *TEntity) *TOutput,
) (*TOutput, *service_errors.ServiceError) {
	entityObj, err := getByIdRepository(ctx, pk)
	if err != nil {
		return nil, service_errors.MapDBErrToServiceErr(
			err, fmt.Sprintf("get %s by id", objectName),
		)
	}

	return toObjectDetailsDTOFunc(entityObj), nil
}

func getOwnerID(
	ctx context.Context,
	pk uint64,
	objectName string,
	getOwnerIdRepository func(ctx context.Context, pk uint64) (uint64, error),
) (uint64, *service_errors.ServiceError) {
	ownerID, err := getOwnerIdRepository(ctx, pk)
	if err != nil {
		return 0, service_errors.MapDBErrToServiceErr(
			err, fmt.Sprintf("get %s owner id", objectName),
		)
	}

	return ownerID, nil
}

func getListOfOuterResourceByFK[TEntity e.TDBEntities, TOutputList generics.OutputListTypes](
	ctx context.Context,
	fk uint64,
	page *d.PaginationQueryParams,
	serviceName string,
	getListOfOuterResourceByFkRepo func(
		ctx context.Context, fk uint64, page *d.PaginationQueryParams,
	) (*d.PagedList[TEntity], error),
	toObjListDtoFunc func([]*TEntity) TOutputList,

) (*generics.PagedList[TOutputList], *service_errors.ServiceError) {

	pagedObjectList, err := getListOfOuterResourceByFkRepo(ctx, fk, page)
	if err != nil {
		return nil, service_errors.MapDBErrToServiceErr(
			err, fmt.Sprintf("%s (id=%d)", serviceName, fk),
		)
	}

	return generics.ToPagedList(pagedObjectList, toObjListDtoFunc), nil
}
