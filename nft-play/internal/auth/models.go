package auth

type AuthorizeRequest struct {
	Address string `json:"address"`
}

type AuthorizeResponse struct {
	Address string `json:"address"`
	Jwt     string `json:"jwt"`
}

type User struct {
	Id      int64  `json:"id" db:"id"`
	Address string `json:"address" db:"address"`
}

type GetUserRequest struct {
	Address string
}

type GetUsersRequest struct {
	Offset int64
	Limit  int64
}

type GetUsersResponse struct {
	Users []User `json:"users"`
}
