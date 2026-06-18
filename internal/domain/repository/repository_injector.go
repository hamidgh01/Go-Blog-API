package repository

type RepositoryInjector interface {
	GetUserRepository() UserRepository
	GetPostRepository() PostRepository
	GetCommentRepository() CommentRepository
	GetTagRepository() TagRepository
	GetLinkRepository() LinkRepository
	GetListRepository() ListRepository
	GetFollowRepository() FollowRepository
	GetLikeRepository() LikeRepository
	GetSavePostRepository() SavePostRepository
	GetSaveListRepository() SaveListRepository
	GetPostTagsRepository() PostTagsRepository
}
