package recipeModel

type RecipeIngredients struct {
	Name       string `bson:"name" json:"name" validate:"nonzero"`
	IsOptional bool   `bson:"isOptional" json:"isOptional"`
}

type Recipe struct {
	Name        string              `bson:"name" json:"name" validate:"nonzero"`
	Ingredients []RecipeIngredients `bson:"ingredients" json:"ingredients" validate:"min=1"`
}
