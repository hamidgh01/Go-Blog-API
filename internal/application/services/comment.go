package services

import (
	"context"
	"database/sql"

	"github.com/hamidgh01/Go-Blog-API/internal/application/service_errors"
	d "github.com/hamidgh01/Go-Blog-API/internal/domain"
	e "github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
	"github.com/hamidgh01/Go-Blog-API/internal/http/generics"
	"github.com/hamidgh01/Go-Blog-API/pkg/constants"
)

type CommentService struct {
	repo repository.CommentRepository
}

func NewCommentService(r repository.CommentRepository) *CommentService {
	return &CommentService{repo: r}
}

func (c *CommentService) Create(
	ctx context.Context, data *dto.CreateCommentRequest,
) (*dto.CommentDetails, *service_errors.ServiceError) {
	//
	userID := ctx.Value(constants.CurrentUserID).(uint64)
	comment := &e.Comment{Content: data.Content, UserID: userID}
	switch data.ParentType {
	case "post":
		comment.PostParentID = sql.Null[uint64]{V: data.ParentID, Valid: true}
	case "comment":
		comment.CommentParentID = sql.Null[uint64]{V: data.ParentID, Valid: true}
	}

	//
	username := ctx.Value(constants.CurrentUserUsername).(string)
	comment.User = &e.User{ID: userID, Username: username}

	//
	return create(ctx, "comment", comment, c.repo.Create, dto.ToCommentDetails)
}

func (c *CommentService) Update(
	ctx context.Context, pk uint64, data *dto.UpdateCommentRequest,
) *service_errors.ServiceError {
	comment := &e.Comment{Content: data.Content}
	return update(ctx, pk, "comment", comment, c.repo.Update)
}

func (c *CommentService) UpdateStatus(
	ctx context.Context, pk uint64, data *dto.UpdateCommentStatusRequest,
) *service_errors.ServiceError {
	err := c.repo.UpdateStatus(ctx, pk, e.CommentStatus(data.Status))
	if err != nil {
		return service_errors.MapDBErrToServiceErr(err, "update comment status")
	}

	return nil
}

func (c *CommentService) Delete(ctx context.Context, pk uint64) *service_errors.ServiceError {
	return delete(ctx, pk, "comment", c.repo.Delete)
}

func (c *CommentService) GetByID(
	ctx context.Context, pk uint64,
) (*dto.CommentDetails, *service_errors.ServiceError) {
	return getByID(ctx, pk, "comment", c.repo.GetByID, dto.ToCommentDetails)
}

// -----------------------------------------------------------------------------
// other sources that has FK to `Comment`

func (c *CommentService) GetReplies(
	ctx context.Context, fk uint64, page *d.PaginationQueryParams,
) (*generics.PagedList[dto.CommentList], *service_errors.ServiceError) {
	return getListOfOuterResourceByFK(
		ctx, fk, page, "get replies of comment", c.repo.GetReplies, dto.ToCommentList,
	)
}

// -----------------------------------------------------------------------------

func (c *CommentService) GetOwnerID(
	ctx context.Context, pk uint64,
) (uint64, *service_errors.ServiceError) {
	return getOwnerID(ctx, pk, "comment", c.repo.GetOwnerID)
}
