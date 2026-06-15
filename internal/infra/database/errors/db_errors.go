package errors

import (
	"errors"

	"github.com/lib/pq"
)

type UniqueViolationError struct {
	message string
}

var _ error = (*UniqueViolationError)(nil)

func newUniqueViolationError(msg string) UniqueViolationError {
	return UniqueViolationError{message: msg}
}

func (s UniqueViolationError) Error() string {
	return s.message
}

var (
	ErrDuplicateEmail    = newUniqueViolationError("email already exists")
	ErrDuplicateUsername = newUniqueViolationError("username already exists")
)

// -----------------------------------------------

type ForeignKeyViolationError struct {
	message string
}

var _ error = (*ForeignKeyViolationError)(nil)

func newForeignKeyViolationError(msg string) ForeignKeyViolationError {
	return ForeignKeyViolationError{message: msg}
}

func (s ForeignKeyViolationError) Error() string {
	return s.message
}

var (
	ErrComments_PostParentIdFkViolation     = newForeignKeyViolationError("invalid request: commenting on non-existent post")
	ErrComments_CommentParentIdFkViolation  = newForeignKeyViolationError("invalid request: replying on non-existent comment")
	ErrFollowsM2M_FollowedIdFkViolation     = newForeignKeyViolationError("invalid request: following non-existent user")
	ErrPostLikesM2M_PostIdFkViolation       = newForeignKeyViolationError("invalid request: liking non-existent post")
	ErrUsersSavedListsM2M_ListIdFkViolation = newForeignKeyViolationError("invalid request: saving non-existent list")
	ErrPostsTagsM2M_PostIdFkViolation       = newForeignKeyViolationError("invalid request: tagging non-existent post")
	ErrPostsTagsM2M_TagIdFkViolation        = newForeignKeyViolationError("invalid request: tagging with non-existent tags")
	ErrSavedPostsM2M_ListIdFkViolation      = newForeignKeyViolationError("invalid request: saving into non-existent list")
	ErrSavedPostsM2M_PostIdFkViolation      = newForeignKeyViolationError("invalid request: saving non-existent post")
)

// -----------------------------------------------

type RecordNotFoundError struct {
	message string
}

var _ error = (*RecordNotFoundError)(nil)

func NewRecordNotFoundError(msg string) RecordNotFoundError {
	return RecordNotFoundError{message: msg}
}

func (s RecordNotFoundError) Error() string {
	return s.message
}

// -----------------------------------------------

type BadInputError struct {
	message string
}

var _ error = (*BadInputError)(nil)

func NewBadInputError(msg string) BadInputError {
	return BadInputError{message: msg}
}

func (s BadInputError) Error() string {
	return s.message
}

// -----------------------------------------------

type UnexpectedDBError struct {
	err     error
	message string
}

var _ error = (*UnexpectedDBError)(nil)

func newUnexpectedDBError(err error) UnexpectedDBError {
	return UnexpectedDBError{err: err, message: err.Error()}
}

func (s UnexpectedDBError) Error() string {
	return s.message
}

func GetDBError(err error) error {

	var pqErr *pq.Error
	if errors.As(err, &pqErr) {

		switch pqErr.Code {

		case "23505": // unique violations
			switch pqErr.Constraint {
			// case "tags_name_key": // ON CONFLICT DO NOTHING
			// 	return ErrDuplicateTagName
			case "users_email_key":
				return ErrDuplicateEmail
			case "users_username_key":
				return ErrDuplicateUsername
			}

		// case "23502": // not null violations (handled by validations)

		case "23514": // check violations
			switch pqErr.Constraint {
			case "cant_follow_yourself":

			}

		case "23503": // foreign key violations
			switch pqErr.Constraint {
			case "comments_postparentid_fkey":
				return ErrComments_PostParentIdFkViolation
			case "comments_commentparentid_fkey":
				return ErrComments_CommentParentIdFkViolation
			case "follows_m2m_followed_fkey":
				return ErrFollowsM2M_FollowedIdFkViolation
			case "post_likes_m2m_post_id_fkey":
				return ErrPostLikesM2M_PostIdFkViolation
			case "users_saved_lists_m2m_list_id_fkey":
				return ErrUsersSavedListsM2M_ListIdFkViolation
			case "posts_tags_m2m_post_id_fkey":
				return ErrPostsTagsM2M_PostIdFkViolation
			case "posts_tags_m2m_tag_id_fkey":
				return ErrPostsTagsM2M_TagIdFkViolation
			case "saved_posts_m2m_list_id_fkey":
				return ErrSavedPostsM2M_ListIdFkViolation
			case "saved_posts_m2m_post_id_fkey":
				return ErrSavedPostsM2M_PostIdFkViolation
			}
		}
	}

	return newUnexpectedDBError(err)
}
