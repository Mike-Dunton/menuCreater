package recipeModel

type RecipeIngredients struct {
	Name       string `json:"name"`
	IsOptional bool   `json:"isOptional"`
}

type Recipe struct {
	Name        string              `json:"name"`
	Ingredients []RecipeIngredients `json:"ingredients"`
}
