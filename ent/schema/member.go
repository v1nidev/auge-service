package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
			&entsql.Annotation{
				Check: "sex in ('M', 'F')",
			},
	}
}

func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("sex"),
		field.String("email").
			MinLen(5).
			MaxLen(255),
	}
}

func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("emergencycontact", EmergencyContact.Type).
			Unique(),
	}
}
