package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemAttributes struct {
	ID          primitive.ObjectID `bson:"id"`
	Name        *string            `json:"name"`
	Description *string            `json:"description"`
	Rarity      *string            `json:"rarity"`
	Damage      *int32             `json:"damage"`
	Defense     *int32             `json:"defense"`
	Vigor       *int32             `json:"vigor"`
	Dexterity   *int32             `json:"dexterity"`
	Empiricism  *int32             `json:"empiricism"`
	Mana        *int32             `json:"mana"`
}
