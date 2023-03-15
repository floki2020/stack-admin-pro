package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Strings("user_name").Comment("登录用户名"),
		field.Strings("password").Comment("密码"),
		field.Int("parent_id").Optional().Comment("父级别ID"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("personals", Personal.Type),
		edge.To("companys", Company.Type),
		edge.To("children", User.Type).From("parent").Field("parent_id").Unique(),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "t_user"},
	}
}
