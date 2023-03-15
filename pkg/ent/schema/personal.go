package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Personal holds the schema definition for the Personal entity.
type Personal struct {
	ent.Schema
}

// Fields of the Personal.
func (Personal) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("个人姓名"),
		field.String("phone").Comment("联系电话"),
		field.Int("user_id").Optional().Comment("关联ID"),
	}
}

// Edges of the Personal.
func (Personal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("personals").Unique().Field("user_id"),
	}
}

func (Personal) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_personal"},
	}
}
