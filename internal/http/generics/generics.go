package generics

import "Go-Blog-API/internal/http/dto"

type CreateRequestTypes interface {
	dto.CreateUserRequest |
		dto.CreatePostRequest |
		dto.CreateCommentRequest |
		dto.CreateListRequest |
		dto.CreateLinkRequest |
		dto.CreateTagsRequest
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
