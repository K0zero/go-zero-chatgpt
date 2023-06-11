// Code generated by goctl. DO NOT EDIT.
package types

type InsertAPIReq struct {
	API string `json:"api"`
}

type InsertAPIResp struct {
	Id     int64  `json:"id"`
	UserId int64  `json:"id"`
	API    string `json:"api"`
}

type DeleteAPIReq struct {
	API string `json:"api"`
}

type DeleteAPIResp struct {
}

type ChatReq struct {
	Question string `json:"question"`
}

type ChatResp struct {
	Answer string `json:"answer"`
}
