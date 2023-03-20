package model

import (
	"context"
	"go-manage/pkg/ent"
)

type defaultUserModel struct {
	*ent.Client
}

type UserInsertReq struct {
	UserName string `json:"user_name"`
}
type UserModel interface {
	InsertOne(ctx context.Context, data *UserInsertReq) error
	DeleteOneById(ctx context.Context, id int) error
}

func NewUserModel(c *ent.Client) UserModel {
	return &defaultUserModel{
		c,
	}
}

func (s *defaultUserModel) InsertOne(ctx context.Context, data *UserInsertReq) error {
	_, err := s.User.Create().SetUserName(data.UserName).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *defaultUserModel) DeleteOneById(ctx context.Context, id int) error {
	err := s.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
