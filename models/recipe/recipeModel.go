package recipeModel

import (
	"gopkg.in/mgo.v2/bson"
)

type RecipeIngredients struct {
	Name       string `bson:"name" json:"name" validate:"nonzero"`
	IsOptional bool   `bson:"isOptional" json:"isOptional"`
}

type Recipe struct {
	ID          bson.ObjectId       `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string              `bson:"name" json:"name" validate:"nonzero"`
	Ingredients []RecipeIngredients `bson:"ingredients" json:"ingredients" validate:"min=1"`
}
