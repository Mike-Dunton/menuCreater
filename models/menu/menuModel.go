package menuModel

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type MenuItem struct {
	Date     time.Time     `bson:"date" json:"name" validate:"nonzero"`
	RecipeID bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
}

type Menu struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	MenuItems []MenuItem    `bson:"menuItems" json:"menuItems" validate:"min=1"`
}
