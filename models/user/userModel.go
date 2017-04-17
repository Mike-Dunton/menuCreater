package userModel

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Email     string        `bson:"email" json:"email" validate:"nonzero"`
	Password  string        `bson:"password" json:"password" validate:"nonzero"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
}
