package dao

import (
	"emperror.dev/errors"

	models "github.com/pgillich/meals-demo/internal/api"
)

func (dbHandler *Handler) CreateTag(tag models.Tag) (models.Tag, error) {
	db := dbHandler.DB.Create(&tag)

	return tag, errors.WrapWithDetails(db.Error, "cannot create tag")
}

func (dbHandler *Handler) GetTags() ([]models.Tag, error) {
	tags := []models.Tag{}
	db := dbHandler.DB.Find(&tags)

	return tags, errors.WrapWithDetails(db.Error, "cannot get tags")
}

func (dbHandler *Handler) fillTags(tags []models.Tag) error {
	if storedTags, err := dbHandler.GetTags(); err != nil {
		return err
	} else if len(storedTags) > 0 {
		return nil
	}

	for _, tag := range tags {
		if _, err := dbHandler.CreateTag(tag); err != nil {
			return err
		}
	}

	return nil
}

func GetDefaultFillTags() []models.Tag {
	return []models.Tag{
		{
			Name: "spicy",
		},
		{
			Name: "vegan",
		},
		{
			Name: "gluten free",
		},
	}
}
