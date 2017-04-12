package menuModel

import "gopkg.in/mgo.v2/bson"

type MenuItem struct {
	Date     string `bson:date`
	RecipeID bson.ObjectId
}

type Menu struct {
	MenuItems []MenuItem `bson:menuItems`
}
