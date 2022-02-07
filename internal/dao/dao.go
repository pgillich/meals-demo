package dao

import (
	"emperror.dev/errors"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres" // import Postgres driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"   // import Sqlite driver

	"github.com/pgillich/meals-demo/configs"
	"github.com/pgillich/meals-demo/internal/models"
)

// Handler is a thin layer over Gorm
type Handler struct {
	DB *gorm.DB
}

func NewHandler(config configs.Options) (*Handler, error) {
	db, err := gorm.Open(config.DbDialect, config.DbDsn)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect to DB")
	}

	if config.DbDebug {
		db = db.Debug()
	}

	dbHandler := &Handler{DB: db}
	if err := dbHandler.prepare(config); err != nil {
		return nil, err
	}

	if config.DbSample {
		if err := dbHandler.fillTags(GetDefaultFillTags()); err != nil {
			return nil, err
		}
		if err := dbHandler.fillIngredients(GetDefaultFillIngredients()); err != nil {
			return nil, err
		}
		tags, err := dbHandler.GetTags()
		if err != nil {
			return nil, err
		}
		ingredients, err := dbHandler.GetIngredients()
		if err != nil {
			return nil, err
		}
		if err := dbHandler.fillMeals(GetDefaultFillMeals(tags, ingredients)); err != nil {
			return nil, err
		}
	}

	return dbHandler, nil
}

func (dbHandler *Handler) Close() {
	dbHandler.DB.Close() //nolint:errcheck,gosec // never mind at exit
}

func (dbHandler *Handler) prepare(config configs.Options) error {
	if dbHandler.DB.Dialect().GetName() == "sqlite3" {
		dbHandler.DB.Exec("PRAGMA foreign_keys = ON")
	}

	if config.DbDebug {
		dbHandler.DB = dbHandler.DB.LogMode(true)
	}

	db := dbHandler.DB.AutoMigrate(&models.Tag{}, &models.Ingredient{}, &models.Meal{})

	return errors.Wrap(db.Error, "cannot update DB schema")
}
