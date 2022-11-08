package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Inventory struct {
	ID             primitive.ObjectID `bson:"id"`
	NomePersonagem *string            `json:"nomePersonagem"`
	Armor          *string            `json:"armor"`
	Helmet         *string            `json:"helmet"`
	Boots          *string            `json:"boots"`
	Pendant        *string            `json:"pendant"`
	Earring        *string            `json:"earring"`
	Stash          *string            `json:"stash"`
	Pants          *string            `json:"pants"`
	Gloves         *string            `json:"gloves"`
	Bag            *string            `json:"nivel"`
}
