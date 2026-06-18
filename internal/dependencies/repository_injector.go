package dependencies

import (
	"database/sql"

	"github.com/hamidgh01/Go-Blog-API/internal/domain/repository"
	"github.com/hamidgh01/Go-Blog-API/internal/infra/database/postgres_repository"
)

type repositoryInjector struct {
	db *sql.DB
}

var _ repository.RepositoryInjector = (*repositoryInjector)(nil)

func NewRepositoryInjector(db *sql.DB) (*repositoryInjector, func()) {
	return &repositoryInjector{db: db}, postgres_repository.CloseAllPreparedStatements
}

func (ri *repositoryInjector) GetUserRepository() repository.UserRepository {
	return postgres_repository.NewUserRepository(ri.db)
}

func (ri *repositoryInjector) GetPostRepository() repository.PostRepository {
	return postgres_repository.NewPostRepository(ri.db)
}

func (ri *repositoryInjector) GetCommentRepository() repository.CommentRepository {
	return postgres_repository.NewCommentRepository(ri.db)
}

func (ri *repositoryInjector) GetTagRepository() repository.TagRepository {
	return postgres_repository.NewTagRepository(ri.db)
}

func (ri *repositoryInjector) GetLinkRepository() repository.LinkRepository {
	return postgres_repository.NewLinkRepository(ri.db)
}

func (ri *repositoryInjector) GetListRepository() repository.ListRepository {
	return postgres_repository.NewListRepository(ri.db)
}

func (ri *repositoryInjector) GetFollowRepository() repository.FollowRepository {
	return postgres_repository.NewFollowRepository(ri.db)
}

func (ri *repositoryInjector) GetLikeRepository() repository.LikeRepository {
	return postgres_repository.NewLikeRepository(ri.db)
}

func (ri *repositoryInjector) GetSavePostRepository() repository.SavePostRepository {
	return postgres_repository.NewSavePostRepository(ri.db)
}

func (ri *repositoryInjector) GetSaveListRepository() repository.SaveListRepository {
	return postgres_repository.NewSaveListRepository(ri.db)
}

func (ri *repositoryInjector) GetPostTagsRepository() repository.PostTagsRepository {
	return postgres_repository.NewPostTagsRepository(ri.db)
}
