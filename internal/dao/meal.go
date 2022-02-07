package dao

import (
	"emperror.dev/errors"

	"github.com/pgillich/meals-demo/internal/models"
)

func (dbHandler *Handler) CreateMeal(meal models.Meal) (models.Meal, error) {
	db := dbHandler.DB.Create(&meal)

	return meal, errors.WrapWithDetails(db.Error, "cannot create meal")
}

func (dbHandler *Handler) UpdateMeal(m models.Meal) (models.Meal, error) {
	meal := &m
	tags := meal.Tags
	ingredients := meal.Ingredients
	meal.Tags = []*models.Tag{}
	meal.Ingredients = []*models.Ingredient{}

	dbA := dbHandler.DB.Model(meal).Association("Tags").Replace(tags)
	if dbA.Error != nil {
		return m, errors.WrapWithDetails(dbA.Error, "cannot update meal")
	}

	dbA = dbHandler.DB.Model(meal).Association("Ingredients").Replace(ingredients)
	if dbA.Error != nil {
		return m, errors.WrapWithDetails(dbA.Error, "cannot update meal")
	}

	db := dbHandler.DB.Model(meal).Select("name", "description", "kcal", "pictureUrl", "price").Updates(meal)
	if db.Error != nil {
		return m, errors.WrapWithDetails(db.Error, "cannot update meal")
	}

	return m, nil
}

func (dbHandler *Handler) DeleteMeal(id int64) error {
	meal := models.Meal{
		ID: id,
	}
	db := dbHandler.DB.Delete(&meal)

	return errors.WrapWithDetails(db.Error, "cannot delete meal")
}

func (dbHandler *Handler) GetMeal(id int64) (*models.Meal, error) {
	meal := models.Meal{
		ID: id,
	}
	db := dbHandler.DB.Preload("Tags").Preload("Ingredients").Find(&meal)

	return &meal, errors.WrapWithDetails(db.Error, "cannot get meal")
}

func (dbHandler *Handler) FindMealsByTag(tagID *int64) ([]*models.Meal, error) {
	meals := []*models.Meal{}
	db := dbHandler.DB.Preload("Tags").Preload("Ingredients").Find(&meals)
	if tagID != nil {
		mealsRemained := []*models.Meal{}
		for _, mea := range meals {
			for _, tag := range mea.Tags {
				if tag.ID == *tagID {
					mealsRemained = append(mealsRemained, mea)

					break
				}
			}
		}
		meals = mealsRemained
	}

	return meals, errors.WrapWithDetails(db.Error, "cannot get meals")
}

func (dbHandler *Handler) fillMeals(meals []*models.Meal) error {
	if storedMeals, err := dbHandler.FindMealsByTag(nil); err != nil {
		return err
	} else if len(storedMeals) > 0 {
		return nil
	}

	for _, meal := range meals {
		if _, err := dbHandler.CreateMeal(*meal); err != nil {
			return err
		}
	}

	return nil
}

func GetDefaultFillMeals(tags []*models.Tag, ingredients []*models.Ingredient) []*models.Meal {
	return []*models.Meal{
		{
			Name:        stringRef("Spicy"),
			Description: "Spicy pizza",
			PictureURL:  "http://a.com",
			Price:       3.25,
			Kcal:        123,
			Ingredients: []*models.Ingredient{
				ingredients[0],
				ingredients[2],
				ingredients[3],
			},
			Tags: []*models.Tag{
				tags[0],
			},
		},
		{
			Name:        stringRef("Vegan"),
			Description: "Vegan pizza",
			PictureURL:  "http://a.com",
			Price:       4.10,
			Kcal:        234,
			Ingredients: []*models.Ingredient{
				ingredients[1],
				ingredients[4],
			},
			Tags: []*models.Tag{
				tags[1],
			},
		},
	}
}

func stringRef(s string) *string {
	return &s
}
