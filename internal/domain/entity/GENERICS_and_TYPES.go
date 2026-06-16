package entity

type TDBEntities interface {
	User | FollowsM2M | Link | Post | PostLikesM2M | Comment |
		Tag | PostTagsM2M | List | SavedPostsM2M | UsersSavedListsM2M
}

type TM2MDBEntities interface {
	FollowsM2M | PostLikesM2M | PostTagsM2M | SavedPostsM2M | UsersSavedListsM2M
}

type CountKey string

const (
	POST_COUNTS      CountKey = "post_counts"
	COMMENT_COUNTS   CountKey = "comment_counts"
	LIST_COUNTS      CountKey = "list_counts"
	LIKE_COUNTS      CountKey = "like_counts"
	LINK_COUNTS      CountKey = "link_counts"
	FOLLOWER_COUNTS  CountKey = "follower_counts"
	FOLLOWING_COUNTS CountKey = "following_counts"
)

type DBEntityWithCountOfReferencedObjects[TEntity TDBEntities] struct {
	Entity       TEntity
	RefObjCounts map[CountKey]int
}
