//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
)

type Member struct {
	ID                 uuid.UUID `sql:"primary_key"`
	Name               string
	Sex                string
	Email              string
	IDMembershipStatus float64
}