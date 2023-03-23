package model

import (
	"context"
	"go-manage/pkg/ent"
	"go-manage/pkg/ent/predicate"
	"go-manage/pkg/ent/user"
)

type defaultUserModel struct {
	*ent.Client
}

type UserInsertReq struct {
	UserName string `json:"user_name"`
}
type UserInfo struct {
	BaseID
	UserName string `json:"user_name"`
	CommonBody
}

type UserParam struct {
	UserName string `json:"user_name"`
	Paging
}
type UserModel interface {
	InsertOne(ctx context.Context, data *UserInsertReq) error
	DeleteOneById(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, data *UserInfo) error
	FindOne(ctx context.Context, id int) (*UserInfo, error)
	FindUserList(ctx context.Context, in UserParam) ([]*UserInfo, int, error)
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

func (s *defaultUserModel) UpdateUser(ctx context.Context, data *UserInfo) error {
	count, err := s.User.Update().SetUserName(data.UserName).Save(ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	return nil
}

func (s *defaultUserModel) FindOne(ctx context.Context, id int) (*UserInfo, error) {
	data, err := s.User.Query().Where(user.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	res := &UserInfo{
		BaseID: BaseID{
			Id: data.ID,
		},
		UserName:   data.UserName,
		CommonBody: CommonBody{},
	}
	return res, nil
}

func (s *defaultUserModel) FindUserList(ctx context.Context, in UserParam) ([]*UserInfo, int, error) {
	var predicates []predicate.User
	if in.UserName != "" {
		predicates = append(predicates, user.UserName(in.UserName))
	}
	users, err := s.User.Query().Where(predicates...).Offset(int((in.Current - 1) * in.PageSize)).Limit(int(in.PageSize)).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	count := len(users)
	var res []*UserInfo

	for _, v := range users {
		data := &UserInfo{
			BaseID: BaseID{
				Id: v.ID,
			},
			UserName:   v.UserName,
			CommonBody: CommonBody{},
		}
		res = append(res, data)
	}
	return res, count, nil
}
