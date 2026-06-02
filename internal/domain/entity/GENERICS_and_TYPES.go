package entity

type TDBEntities interface {
	User | FollowsM2M | Link | Post | PostLikesM2M | Comment |
		Tag | PostTagsM2M | List | SavedPostsM2M | UsersSavedListsM2M
}

type ReferencedObjectKey string

const (
	LINKS       ReferencedObjectKey = "links"
	POSTS       ReferencedObjectKey = "posts"
	TAGS        ReferencedObjectKey = "tags"
	POSTS_TAGS  ReferencedObjectKey = "posts_tags_m2m"
	COMMENTS    ReferencedObjectKey = "comments"
	LISTS       ReferencedObjectKey = "lists"
	SAVED_POSTS ReferencedObjectKey = "saved_posts_m2m"
	SAVED_LISTS ReferencedObjectKey = "users_saved_list_m2m"
	LIKES       ReferencedObjectKey = "post_likes_m2m"
)

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
