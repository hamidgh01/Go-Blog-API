package entity

type TDBEntities interface {
	User | FollowsM2M | Link | Post | PostLikesM2M | Comment |
		Tag | PostTagsM2M | List | SavedPostsM2M | UsersSavedListsM2M
}
