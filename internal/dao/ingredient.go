package dao

import (
	"emperror.dev/errors"

	"github.com/pgillich/meals-demo/internal/models"
)

func (dbHandler *Handler) CreateIngredient(ingredient models.Ingredient) (models.Ingredient, error) {
	db := dbHandler.DB.Create(&ingredient)

	return ingredient, errors.WrapWithDetails(db.Error, "cannot create ingredient")
}

func (dbHandler *Handler) GetIngredients() ([]*models.Ingredient, error) {
	ingredients := []*models.Ingredient{}
	db := dbHandler.DB.Find(&ingredients)

	return ingredients, errors.WrapWithDetails(db.Error, "cannot get ingredients")
}

func (dbHandler *Handler) fillIngredients(ingredients []*models.Ingredient) error {
	if storedIngredients, err := dbHandler.GetIngredients(); err != nil {
		return err
	} else if len(storedIngredients) > 0 {
		return nil
	}

	for _, ingredient := range ingredients {
		if _, err := dbHandler.CreateIngredient(*ingredient); err != nil {
			return err
		}
	}

	return nil
}

func GetDefaultFillIngredients() []*models.Ingredient {
	return []*models.Ingredient{
		{
			Name:        "tomato sauce",
			Description: "Tomato sauce",
		},
		{
			Name:        "sour cream sauce",
			Description: "Sour cream sauce",
		},
		{
			Name:        "bacon",
			Description: "Bacon",
		},
		{
			Name:        "salami",
			Description: "Salami",
		},
		{
			Name:        "mozzarella",
			Description: "Mozzarella",
		},
		{
			Name:        "tomato",
			Description: "Tomato",
		},
		{
			Name:        "onion",
			Description: "Onion",
		},
	}
}
