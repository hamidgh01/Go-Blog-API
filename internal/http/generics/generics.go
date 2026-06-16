package generics

import (
	"github.com/hamidgh01/Go-Blog-API/internal/domain"
	"github.com/hamidgh01/Go-Blog-API/internal/domain/entity"
	"github.com/hamidgh01/Go-Blog-API/internal/http/dto"
)

type CreateRequestTypes interface {
	dto.CreateUserRequest |
		dto.CreatePostRequest |
		dto.CreateCommentRequest |
		dto.CreateListRequest |
		dto.CreateLinkRequest |
		dto.CreateTagsRequest
}

type CreateOrDeleteM2MRelationshipRequestTypes interface {
	dto.AssociatePostWithTagsRequest |
		dto.DissociatePostWithTagsRequest |
		dto.SavePostRequest |
		dto.UnsavePostRequest
}

type UpdateRequestTypes interface {
	dto.UpdateUsernameRequest |
		dto.UpdateEmailRequest |
		dto.UpdateBioRequest |
		dto.UpdatePasswordRequest |
		dto.UpdateEnabledRequest |
		dto.ResetPasswordRequest |

		dto.UpdatePostRequest |
		dto.UpdatePostStatusRequest |
		dto.UpdatePostPrivacyRequest |

		dto.UpdateCommentRequest |
		dto.UpdateCommentStatusRequest |

		dto.UpdateLinkRequest |

		dto.UpdateListRequest |
		dto.UpdateListPrivacyRequest
}

type OutputTypes interface {
	dto.UserDetails |
		dto.PostDetails |
		dto.CommentDetails |
		dto.ListDetails |
		dto.TagDetails |
		dto.LinkDetails
}

type OutputWithRefObjCountsTypes interface {
	dto.UserDetailsWithCountOfReferencedObjects |
		dto.PostDetailsWithCountOfReferencedObjects |
		dto.CommentDetailsWithRepliesCount |
		dto.ListDetailsWithCountOfReferencedObjects
}

type OutputListTypes interface {
	dto.UsersList |
		dto.PostsList |
		dto.CommentList |
		dto.ListsList |
		dto.LinksList |
		dto.TagsList |
		dto.FollowersList | dto.FollowingsList |
		dto.UsersWhoLikedAPost | dto.PostsLikedByAUser |
		dto.ListOfSavedPosts | dto.ListOfListsAPostSavedIn |
		dto.ListOfSavedLists | dto.ListOfUsersWhoSavedAList
}

type PagedList[TOutputList OutputListTypes] struct {
	Page            int         `json:"page"`
	Size            int         `json:"size"`
	TotalRows       int         `json:"total_rows"`
	TotalPages      int         `json:"total_pages"`
	HasPreviousPage bool        `json:"has_previous_page"`
	PreviousPage    int         `json:"previous_page,omitempty"`
	HasNextPage     bool        `json:"has_next_page"`
	NextPage        int         `json:"next_page,omitempty"`
	Items           TOutputList `json:"items"`
}

func ToPagedList[TOutputList OutputListTypes, TEntity entity.TDBEntities](
	entityPagedList *domain.PagedList[TEntity],
	toObjectListDTOFunc func(entities []*TEntity) TOutputList,
) *PagedList[TOutputList] {
	return &PagedList[TOutputList]{
		Page:            entityPagedList.Page,
		Size:            entityPagedList.Size,
		TotalRows:       entityPagedList.TotalRows,
		TotalPages:      entityPagedList.TotalPages,
		HasPreviousPage: entityPagedList.HasPreviousPage,
		PreviousPage:    entityPagedList.PreviousPage,
		HasNextPage:     entityPagedList.HasNextPage,
		NextPage:        entityPagedList.NextPage,
		Items:           toObjectListDTOFunc(entityPagedList.Items),
	}
}
