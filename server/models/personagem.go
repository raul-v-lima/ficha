package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Personagem struct {
	ID           primitive.ObjectID `bson:"id"`
	Nome         *string            `json:"nome"`
	Forca        *int32             `json:"forca"`
	Armadura     *int32             `json:"armadura"`
	Resistencia  *int32             `json:"resistencia"`
	Destreza     *int32             `json:"destreza"`
	Inteligencia *int32             `json:"inteligencia"`
	Mana         *int32             `json:"mana"`
	Xp           *int32             `json:"xp"`
	Nivel        *int32             `json:"nivel"`
}
