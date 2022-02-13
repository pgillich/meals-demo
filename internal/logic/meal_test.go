package logic

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/pgillich/meals-demo/configs"
	"github.com/pgillich/meals-demo/internal"
	"github.com/pgillich/meals-demo/internal/dao"
	"github.com/pgillich/meals-demo/internal/models"
	"github.com/pgillich/meals-demo/internal/restapi"
)

type MealTestSuite struct {
	suite.Suite
}

func TestMealTestSuite(t *testing.T) {
	os.Args = os.Args[:1]
	suite.Run(t, new(InfoTestSuite))
}

func (s *InfoTestSuite) buildServerMeal(options *configs.Options) *restapi.Server {
	return internal.BuildServer(options, SetUserAPI, SetMealAPI)
}

func (s *InfoTestSuite) getJwtToken(server *restapi.Server) string {
	login := &models.LoginInfo{
		Email:    stringRef("yoda@star.wars"),
		Password: stringRef("master"),
	}
	body, err := json.Marshal(login)
	s.NoError(err, "Token")
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", configs.TestingBasePath+"/login", bytes.NewReader(body))
	r.Header.Add("Content-Type", "application/json")

	server.GetHandler().ServeHTTP(w, r)

	body = w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	tokenCreated := &models.LoginSuccess{}
	err = json.Unmarshal(body, tokenCreated)
	s.NoError(err, "Body")

	return tokenCreated.Token
}

func (s *InfoTestSuite) TestMealGetTags() {
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     ":memory:",
		DbSample:  true,
		DbDebug:   false,
		JwtKey:    "5678",
	}
	server := s.buildServerMeal(&options)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", configs.TestingBasePath+"/tags", nil)

	server.GetHandler().ServeHTTP(w, r)

	defaultTags := guessTagIDs(dao.GetDefaultFillTags())
	body := w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	tags := []*models.Tag{}
	err := json.Unmarshal(body, &tags)
	s.NoError(err, "Body")
	s.ElementsMatch(defaultTags, tags)
}

func (s *InfoTestSuite) TestMealGetIngredients() {
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     ":memory:",
		DbSample:  true,
		DbDebug:   false,
		JwtKey:    "5678",
	}
	server := s.buildServerMeal(&options)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", configs.TestingBasePath+"/ingredients", nil)

	server.GetHandler().ServeHTTP(w, r)

	defaultIngredients := guessIngredientIDs(dao.GetDefaultFillIngredients())
	body := w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	ingredients := []*models.Ingredient{}
	err := json.Unmarshal(body, &ingredients)
	s.NoError(err, "Body")
	s.ElementsMatch(defaultIngredients, ingredients)
}

func (s *InfoTestSuite) TestMealFindMeals() {
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     ":memory:",
		DbSample:  true,
		DbDebug:   false,
		JwtKey:    "5678",
	}
	server := s.buildServerMeal(&options)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", configs.TestingBasePath+"/meal/findByTag", nil)

	server.GetHandler().ServeHTTP(w, r)

	defaultMeals := guessMealIDs(dao.GetDefaultFillMeals(
		guessTagIDs(dao.GetDefaultFillTags()),
		guessIngredientIDs(dao.GetDefaultFillIngredients()),
	))
	body := w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	// fmt.Println(string(body))
	meals := []*models.Meal{}
	err := json.Unmarshal(body, &meals)
	s.NoError(err, "Body")
	s.ElementsMatch(defaultMeals, meals)
}

func (s *InfoTestSuite) TestMealFindMealsByTags() {
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     ":memory:",
		DbSample:  true,
		DbDebug:   false,
		JwtKey:    "5678",
	}
	server := s.buildServerMeal(&options)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", configs.TestingBasePath+"/meal/findByTag?tag=1", nil)

	server.GetHandler().ServeHTTP(w, r)

	body := w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	//fmt.Println(string(body))
	meals := []*models.Meal{}
	err := json.Unmarshal(body, &meals)
	s.NoError(err, "Body")
}

func (s *InfoTestSuite) TestMealDeleteMeal() {
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     ":memory:",
		DbSample:  true,
		DbDebug:   false,
		JwtKey:    "5678",
	}
	server := s.buildServerMeal(&options)
	token := s.getJwtToken(server)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", configs.TestingBasePath+"/meal/1", nil)
	r.Header.Add("Authorization", "Bearer "+token)

	server.GetHandler().ServeHTTP(w, r)

	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+w.Body.String())
	//body := w.Body.Bytes()
	//fmt.Println(string(body))
}

func (s *InfoTestSuite) TestMealUpdateMeal() {
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     ":memory:",
		DbSample:  true,
		DbDebug:   false,
		JwtKey:    "5678",
	}
	server := s.buildServerMeal(&options)
	token := s.getJwtToken(server)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", configs.TestingBasePath+"/meal/1", nil)

	server.GetHandler().ServeHTTP(w, r)

	body := w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	//fmt.Println(string(body))
	meal := &models.Meal{}
	err := json.Unmarshal(body, meal)
	s.NoError(err, "Body")

	meal.Description = "TEST"
	meal.Ingredients = meal.Ingredients[:1]
	body, err = json.Marshal(meal)
	s.NoError(err, "Body")
	w = httptest.NewRecorder()
	r = httptest.NewRequest("PUT", configs.TestingBasePath+"/meal/0", bytes.NewReader(body))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+token)

	server.GetHandler().ServeHTTP(w, r)

	body = w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	//fmt.Println(string(body))
	mealUpdated := &models.Meal{}
	err = json.Unmarshal(body, mealUpdated)
	s.NoError(err, "Body")
	s.Equal("TEST", meal.Description, "Update")

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", configs.TestingBasePath+"/meal/1", nil)

	server.GetHandler().ServeHTTP(w, r)

	body = w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	//fmt.Println(string(body))
	mealLatest := &models.Meal{}
	err = json.Unmarshal(body, mealLatest)
	s.NoError(err, "Body")
	s.Equal("TEST", mealLatest.Description, "Update Description")
	s.ElementsMatch(meal.Tags, mealLatest.Tags, "Update Tags")
	s.ElementsMatch(meal.Ingredients, mealLatest.Ingredients, "Update Ingredients")
}

func (s *InfoTestSuite) TestMealCreateMeal() {
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     ":memory:",
		DbSample:  true,
		DbDebug:   true,
		JwtKey:    "5678",
	}
	server := s.buildServerMeal(&options)
	token := s.getJwtToken(server)
	tags := guessTagIDs(dao.GetDefaultFillTags())
	ingredients := guessIngredientIDs(dao.GetDefaultFillIngredients())
	meal := &models.Meal{
		Name:        stringRef("Vegan"),
		Description: "Vegan pizza",
		PictureURL:  "http://a.com",
		Price:       4.10,
		Kcal:        234,
		Ingredients: []*models.Ingredient{
			ingredients[0],
			ingredients[5],
			ingredients[6],
		},
		Tags: []*models.Tag{
			tags[2],
		},
	}
	body, err := json.Marshal(meal)
	s.NoError(err, "Body")
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", configs.TestingBasePath+"/meal/0", bytes.NewReader(body))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+token)

	server.GetHandler().ServeHTTP(w, r)

	body = w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	//fmt.Println(string(body))
	mealCreated := &models.Meal{}
	err = json.Unmarshal(body, mealCreated)
	s.NoError(err, "Body")
}

func guessTagIDs(tags []*models.Tag) []*models.Tag {
	for t, tag := range tags {
		tag.ID = int64(t + 1)
	}

	return tags
}

func guessIngredientIDs(ingredients []*models.Ingredient) []*models.Ingredient {
	for i, tag := range ingredients {
		tag.ID = int64(i + 1)
	}

	return ingredients
}

func guessMealIDs(meals []*models.Meal) []*models.Meal {
	for m, mea := range meals {
		mea.ID = int64(m + 1)
	}

	return meals
}

func stringRef(s string) *string {
	return &s
}
