package logic

import (
	"context"

	"github.com/pgillich/meals-demo/internal/api"
)

func (fs *FoodStore) GetTags(ctx context.Context, request api.GetTagsRequestObject) (api.GetTagsResponseObject, error) {
	if tags, err := fs.dbHandler.GetTags(); err != nil {
		return api.GetTags500JSONResponse{Message: err.Error()}, nil //nolint:nilerr // in another field
	} else {
		return api.GetTags200JSONResponse(tags), nil
	}
}

func (fs *FoodStore) GetIngredients(
	ctx context.Context, request api.GetIngredientsRequestObject,
) (api.GetIngredientsResponseObject, error) {
	if ingredients, err := fs.dbHandler.GetIngredients(); err != nil {
		return api.GetIngredients500JSONResponse{Message: err.Error()}, nil //nolint:nilerr // in another field
	} else {
		return api.GetIngredients200JSONResponse(ingredients), nil
	}
}

func (fs *FoodStore) CreateMeal(ctx context.Context, request api.CreateMealRequestObject) (api.CreateMealResponseObject, error) {
	if request.Body == nil {
		return api.CreateMeal500JSONResponse{Message: "meal is required"}, nil
	}

	if m, err := fs.dbHandler.CreateMeal(*request.Body); err != nil {
		return api.CreateMeal500JSONResponse{Message: "meal is required"}, nil //nolint:nilerr // in another field
	} else {
		return api.CreateMeal200JSONResponse(m), nil
	}
}

func (fs *FoodStore) UpdateMeal(ctx context.Context, request api.UpdateMealRequestObject) (api.UpdateMealResponseObject, error) {
	if request.Body == nil {
		return api.UpdateMeal500JSONResponse{Message: "meal is required"}, nil
	}

	if m, err := fs.dbHandler.UpdateMeal(*request.Body); err != nil {
		return api.UpdateMeal500JSONResponse{Message: err.Error()}, nil //nolint:nilerr // in another field
	} else {
		return api.UpdateMeal200JSONResponse(m), nil
	}
}

func (fs *FoodStore) DeleteMeal(ctx context.Context, request api.DeleteMealRequestObject) (api.DeleteMealResponseObject, error) {
	if err := fs.dbHandler.DeleteMeal(request.Id); err != nil {
		return api.DeleteMeal500JSONResponse{Message: err.Error()}, nil //nolint:nilerr // in another field
	}

	return api.DeleteMeal200Response{}, nil
}

func (fs *FoodStore) GetMealById( //nolint:golint,revive // generated func name
	ctx context.Context, request api.GetMealByIdRequestObject,
) (api.GetMealByIdResponseObject, error) {
	if mea, err := fs.dbHandler.GetMeal(request.Id); err != nil {
		return api.GetMealById500JSONResponse{Message: err.Error()}, nil //nolint:nilerr // in another field
	} else {
		return api.GetMealById200JSONResponse(*mea), nil
	}
}

func (fs *FoodStore) FindMealsByTag(
	ctx context.Context, request api.FindMealsByTagRequestObject,
) (api.FindMealsByTagResponseObject, error) {
	if meals, err := fs.dbHandler.FindMealsByTag(request.Params.Tag); err != nil {
		return api.FindMealsByTag500JSONResponse{Message: err.Error()}, nil //nolint:nilerr // in another field
	} else {
		return api.FindMealsByTag200JSONResponse(meals), nil
	}
}
