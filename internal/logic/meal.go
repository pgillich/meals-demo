package logic

import (
	"log"

	"github.com/go-openapi/runtime/middleware"

	"github.com/pgillich/meals-demo/configs"
	"github.com/pgillich/meals-demo/internal/dao"
	"github.com/pgillich/meals-demo/internal/models"
	"github.com/pgillich/meals-demo/internal/restapi/operations"
	"github.com/pgillich/meals-demo/internal/restapi/operations/meal"
)

func SetMealAPI(config configs.Options, api *operations.OpenAPIFoodstoreAPI) {
	mealAPI := &MealAPI{}
	var err error

	mealAPI.dbHandler, err = dao.NewHandler(config)
	if err != nil {
		log.Fatal(err)
	}

	api.MealGetTagsHandler = meal.GetTagsHandlerFunc(mealAPI.GetTags)
	api.MealGetIngredientsHandler = meal.GetIngredientsHandlerFunc(mealAPI.GetIngredients)
	api.MealCreateMealHandler = meal.CreateMealHandlerFunc(mealAPI.CreateMeal)
	api.MealUpdateMealHandler = meal.UpdateMealHandlerFunc(mealAPI.UpdateMeal)
	api.MealDeleteMealHandler = meal.DeleteMealHandlerFunc(mealAPI.DeleteMeal)
	api.MealGetMealByIDHandler = meal.GetMealByIDHandlerFunc(mealAPI.GetMealByID)
	api.MealFindMealsByTagHandler = meal.FindMealsByTagHandlerFunc(mealAPI.FindMealsByTag)
}

type MealAPI struct {
	dbHandler *dao.Handler
}

func (mealAPI *MealAPI) GetTags(params meal.GetTagsParams) middleware.Responder {
	if tags, err := mealAPI.dbHandler.GetTags(); err != nil {
		return meal.NewGetTagsInternalServerError().WithPayload(&models.APIError{
			Message: err.Error(),
		})
	} else {
		return meal.NewGetTagsOK().WithPayload(tags)
	}
}

func (mealAPI *MealAPI) GetIngredients(params meal.GetIngredientsParams) middleware.Responder {
	if ingredients, err := mealAPI.dbHandler.GetIngredients(); err != nil {
		return meal.NewGetIngredientsInternalServerError().WithPayload(&models.APIError{
			Message: err.Error(),
		})
	} else {
		return meal.NewGetIngredientsOK().WithPayload(ingredients)
	}
}

func (mealAPI *MealAPI) CreateMeal(params meal.CreateMealParams, user *models.User) middleware.Responder {
	if params.Body == nil {
		return meal.NewCreateMealInternalServerError().WithPayload(&models.APIError{
			Message: "meal is required",
		})
	}

	if m, err := mealAPI.dbHandler.CreateMeal(*params.Body); err != nil {
		return meal.NewCreateMealInternalServerError().WithPayload(&models.APIError{
			Message: err.Error(),
		})
	} else {
		return meal.NewCreateMealOK().WithPayload(&m)
	}
}

func (mealAPI *MealAPI) UpdateMeal(params meal.UpdateMealParams, user *models.User) middleware.Responder {
	if params.Body == nil {
		return meal.NewUpdateMealInternalServerError().WithPayload(&models.APIError{
			Message: "meal is required",
		})
	}

	if m, err := mealAPI.dbHandler.UpdateMeal(*params.Body); err != nil {
		return meal.NewUpdateMealInternalServerError().WithPayload(&models.APIError{
			Message: err.Error(),
		})
	} else {
		return meal.NewUpdateMealOK().WithPayload(&m)
	}
}

func (mealAPI *MealAPI) DeleteMeal(params meal.DeleteMealParams, user *models.User) middleware.Responder {
	if err := mealAPI.dbHandler.DeleteMeal(params.ID); err != nil {
		return meal.NewDeleteMealInternalServerError().WithPayload(&models.APIError{
			Message: err.Error(),
		})
	}

	return meal.NewDeleteMealOK()
}

func (mealAPI *MealAPI) GetMealByID(params meal.GetMealByIDParams) middleware.Responder {
	if mea, err := mealAPI.dbHandler.GetMeal(params.ID); err != nil {
		return meal.NewGetMealByIDInternalServerError().WithPayload(&models.APIError{
			Message: err.Error(),
		})
	} else {
		return meal.NewGetMealByIDOK().WithPayload(mea)
	}
}

func (mealAPI *MealAPI) FindMealsByTag(params meal.FindMealsByTagParams) middleware.Responder {
	if meals, err := mealAPI.dbHandler.FindMealsByTag(params.Tag); err != nil {
		return meal.NewFindMealsByTagInternalServerError().WithPayload(&models.APIError{
			Message: err.Error(),
		})
	} else {
		return meal.NewFindMealsByTagOK().WithPayload(meals)
	}
}
