package user

import (
	"apiserver/model"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Sex      int    `json:"sex"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Birthday string `json:"birthday" `
}
type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `form:"username"`
	Offset   int    `form:"offset"`
	Limit    int    `form:"limit"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
