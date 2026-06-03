package dto

import "Go-Blog-API/internal/domain/entity"

type UsersWhoLikedAPost struct {
	Users UsersList `json:"users"`
}

func ToUsersWhoLikedAPost(users []*entity.User) *UsersWhoLikedAPost {
	return &UsersWhoLikedAPost{Users: ToUsersList(users)}
}

type PostsLikedByAUser struct {
	Posts PostsList `json:"posts"`
}

func ToPostsLikedByAUser(posts []*entity.Post) *PostsLikedByAUser {
	return &PostsLikedByAUser{Posts: ToPostsList(posts)}
}
