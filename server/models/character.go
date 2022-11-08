package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Character struct {
	ID         primitive.ObjectID `bson:"id"`
	Name       *string            `json:"name"`
	Vigor      *int32             `json:"vigor"`
	Empiricism *int32             `json:"empiricism"`
	Dexterity  *int32             `json:"dexterity"`
	Mana       *int32             `json:"mana"`
	Xp         *int32             `json:"xp"`
	Level      *int32             `json:"level"`
}
