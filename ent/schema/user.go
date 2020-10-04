package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("firstname"),
		field.String("lastname"),
		field.String("username").Unique(),
		field.String("password").Sensitive(),
		field.String("address").Optional(),
		field.String("city").Optional(),
		field.String("state").Optional(),
		field.String("zipcode").Optional(),
		field.String("email").Optional(),
		field.String("phone").Optional(),
		field.String("photo").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type),
	}
}
