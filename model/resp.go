package model

type OperationResp struct {
	Code uint   `json:"code"`
	Msg  string `json:"msg"`
}

type UsersResp struct {
	Code  uint   `json:"code"`
	Users []User `json:"users"`
}

type TokenResp struct {
	Code  uint   `json:"code"`
	Token string `json:"token"`
	Role  int    `json:"role"`
}

type UrlRersp struct {
	Code uint   `json:"code"`
	Url  string `json:"url"`
}