package model

import "time"

type BaseID struct {
	Id int `json:"id"`
}

type UUIDBase struct {
	UUID string `json:"uuid"`
}

type CommonBody struct {
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
}

type Paging struct {
	Current  int `json:"current"`
	PageSize int `json:"page_size"`
}
