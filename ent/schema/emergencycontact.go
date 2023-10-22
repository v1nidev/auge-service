package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// EmergencyContact holds the schema definition for the EmergencyContact entity.
type EmergencyContact struct {
	ent.Schema
}

func (EmergencyContact) Fields() []ent.Field {
	return []ent.Field{
		field.String("phone_number").
			MaxLen(20),
		field.String("contact_name"),
	}
}

func (EmergencyContact) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `Member`
		// and reference it to the "member" edge (in Member schema)
		// explicitly using the `Ref` method.
		edge.From("referrer", Member.Type).
				Ref("emergencycontact").
				// setting the edge to unique, ensure
				// that a car can have only one owner.
				Unique().
				// We add the "Required" method to the builder
				// to make this edge required on entity creation.
				// i.e. EmergencyContact cannot be created without its referrer.
				Required(),
	}
}
