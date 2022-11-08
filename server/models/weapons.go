package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Weapons struct {
	ID              primitive.ObjectID `bson:"id"`
	Name            *string            `json:"name"`
	Type            *string            `json:"type"`
	DamageType      *string            `json:"damageType"`
	Rarity          *string            `json:"boots"`
	Pendant         *string            `json:"Rarity"`
	RequiredLevel   *string            `json:"requiredLevel"`
	Attributes      *string            `json:"Attributes"`
	MagicAttributes *string            `json:"magicAttributes"`
}
