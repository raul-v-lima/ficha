package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Environment struct {
	ID                primitive.ObjectID `bson:"id"`
	Name              *string            `json:"name"`
	Vegetation        *string            `json:"vegetation"`
	Terrain           *string            `json:"terrain"`
	Creatures         *string            `json:"creatures"` // beasts, animals,monsters
	EndemicCreatures  *string            `json:"endemicCreatures"`
	Climate           *string            `json:"climate"`
	EndemicVegetation *string            `json:"endemicVegetation"` // rare herbs,fruits,wood...
	Landscape         *string            `json:"landscape"`         // meadows,mountainous,
	Buildings         *string            `json:"buildings"`         // churches, caves, shrines, ruins...
	WaterBodies       *string            `json:"waterBodies"`       // lakes, swamps,rivers, oceans , bays...
	ExoticAssets      *string            `json:"exoticAssets"`      // vulcano X , river Y, crater Z , portals X ....
	Decoys            *string            `json:"decoys"`            // loots, jewels,rings, trinquets,weapons, treasures
	ContextCharacters *string            `json:"contextCharacters"` // local NPCs , enemies, entities
}
