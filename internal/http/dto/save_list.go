package dto

import "Go-Blog-API/internal/domain/entity"

type ListOfSavedLists struct {
	Lists ListsList `json:"lists"`
}

func ToListOfSavedLists(lists []*entity.List) *ListOfSavedLists {
	return &ListOfSavedLists{Lists: ToListsList(lists)}
}

type ListOfUsersWhoSavedAList struct {
	Users UsersList `json:"users"`
}

func ToListOfUsersWhoSavedAList(users []*entity.User) *ListOfUsersWhoSavedAList {
	return &ListOfUsersWhoSavedAList{Users: ToUsersList(users)}
}
