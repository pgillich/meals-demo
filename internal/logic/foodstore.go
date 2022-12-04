package logic

import (
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/pgillich/meals-demo/configs"
	"github.com/pgillich/meals-demo/internal/api"
	"github.com/pgillich/meals-demo/internal/dao"
)

type FoodStore struct {
	dbHandler *dao.Handler

	jwtIssuer    string
	jwtKey       string
	jwtExpireSec time.Duration
	method       jwt.SigningMethod
}

func NewFoodStore(config configs.Options) api.StrictServerInterface {
	dbHandler, err := dao.NewHandler(config)
	if err != nil {
		log.Fatal(err)
	}

	var jwtExpireSec int
	if config.JwtExpireSec != "" {
		if jwtExpireSec, err = strconv.Atoi(config.JwtExpireSec); err != nil {
			log.Fatal(err)
		}
	}
	if jwtExpireSec <= 0 {
		jwtExpireSec = 60 * 60
	}

	return &FoodStore{
		dbHandler: dbHandler,

		jwtIssuer:    "foodstore",
		jwtKey:       config.JwtKey,
		jwtExpireSec: time.Duration(jwtExpireSec) * time.Second,
		method:       jwt.SigningMethodHS256,
	}

	/*
		api.JWTAuth = userAPI.ValidateHeader
		api.APIKeyAuthenticator = security.APIKeyAuth
		api.UserLoginHandler = user.LoginHandlerFunc(userAPI.Login)
	*/
}

/*
// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all ingredients
	// (GET /ingredients)
	GetIngredients(ctx context.Context, request GetIngredientsRequestObject) (GetIngredientsResponseObject, error)
	// Liveness status for orchestrator
	// (GET /livez)
	GetLivez(ctx context.Context, request GetLivezRequestObject) (GetLivezResponseObject, error)

	// (POST /login)
	Login(ctx context.Context, request LoginRequestObject) (LoginResponseObject, error)
	// Finds Meals by tag
	// (GET /meal/findByTag)
	FindMealsByTag(ctx context.Context, request FindMealsByTagRequestObject) (FindMealsByTagResponseObject, error)
	// Deletes a meal
	// (DELETE /meal/{id})
	DeleteMeal(ctx context.Context, request DeleteMealRequestObject) (DeleteMealResponseObject, error)
	// Find meal by ID
	// (GET /meal/{id})
	GetMealById(ctx context.Context, request GetMealByIdRequestObject) (GetMealByIdResponseObject, error)
	// Create a new meal
	// (POST /meal/{id})
	CreateMeal(ctx context.Context, request CreateMealRequestObject) (CreateMealResponseObject, error)
	// Update an existing meal
	// (PUT /meal/{id})
	UpdateMeal(ctx context.Context, request UpdateMealRequestObject) (UpdateMealResponseObject, error)
	// Get all tags
	// (GET /tags)
	GetTags(ctx context.Context, request GetTagsRequestObject) (GetTagsResponseObject, error)
	// Version
	// (GET /version)
	GetVersion(ctx context.Context, request GetVersionRequestObject) (GetVersionResponseObject, error)
}


*/
