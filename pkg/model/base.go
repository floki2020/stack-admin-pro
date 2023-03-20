package model

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"go-manage/pkg/ent"
)

type defaultBaseModel struct {
	*ent.Client
}

type BaseModel interface {
	Init(ctx context.Context) error
}

func NewBaseModel(c *ent.Client) BaseModel {
	return &defaultBaseModel{
		c,
	}
}

func (s *defaultBaseModel) Init(ctx context.Context) error {
	err := s.Schema.Create(ctx, schema.WithForeignKeys(false), schema.WithDropColumn(true),
		schema.WithDropIndex(true))
	if err != nil {
		return err
	}
	return nil
}
