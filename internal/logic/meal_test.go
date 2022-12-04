package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/pgillich/meals-demo/configs"
	"github.com/pgillich/meals-demo/internal"
	"github.com/pgillich/meals-demo/internal/api"
	"github.com/pgillich/meals-demo/internal/dao"
)

type MealTestSuite struct {
	suite.Suite
}

func TestMealTestSuite(t *testing.T) {
	os.Args = os.Args[:1]
	suite.Run(t, new(InfoTestSuite))
}

func (s *InfoTestSuite) buildServerMeal(options *configs.Options) *http.Server {
	return internal.BuildServer(options, NewFoodStore)
}

func (s *InfoTestSuite) getJwtToken(server *http.Server) string {
	login := &api.LoginInfo{
		Email:    "yoda@star.wars",
		Password: "master",
	}
	body, err := json.Marshal(login)
	s.NoError(err, "Token")
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", configs.TestingBasePath+"/login", bytes.NewReader(body))
	r.Header.Add("Content-Type", "application/json")

	server.Handler.ServeHTTP(w, r)
	defer server.Close()

	body = w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	tokenCreated := &api.LoginSuccess{}
	err = json.Unmarshal(body, tokenCreated)
	s.NoError(err, "Body")

	return tokenCreated.Token
}

func (s *InfoTestSuite) TestMealGetTags() {
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     fmt.Sprintf("file:memdb%s%d?mode=memory&cache=shared&_pragma=foreign_keys(1)", regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(s.T().Name(), "_"), os.Getpid()),
		DbSample:  true,
		DbDebug:   false,
		JwtKey:    "5678",
	}
	server := s.buildServerMeal(&options)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", configs.TestingBasePath+"/tags", nil)

	server.Handler.ServeHTTP(w, r)
	defer server.Close()

	defaultTags := guessTagIDs(dao.GetDefaultFillTags())
	body := w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	tags := []api.Tag{}
	err := json.Unmarshal(body, &tags)
	s.NoError(err, "Body")
	s.ElementsMatch(defaultTags, tags)
}

func (s *InfoTestSuite) TestMealGetIngredients() {
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     fmt.Sprintf("file:memdb%s%d?mode=memory&cache=shared&_pragma=foreign_keys(1)", regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(s.T().Name(), "_"), os.Getpid()),
		DbSample:  true,
		DbDebug:   false,
		JwtKey:    "5678",
	}
	server := s.buildServerMeal(&options)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", configs.TestingBasePath+"/ingredients", nil)

	server.Handler.ServeHTTP(w, r)
	defer server.Close()

	defaultIngredients := guessIngredientIDs(dao.GetDefaultFillIngredients())
	body := w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	ingredients := []api.Ingredient{}
	err := json.Unmarshal(body, &ingredients)
	s.NoError(err, "Body")
	s.ElementsMatch(defaultIngredients, ingredients)
}

func (s *InfoTestSuite) TestMealFindMeals() {
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     fmt.Sprintf("file:memdb%s%d?mode=memory&cache=shared&_pragma=foreign_keys(1)", regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(s.T().Name(), "_"), os.Getpid()),
		DbSample:  true,
		DbDebug:   false,
		JwtKey:    "5678",
	}
	server := s.buildServerMeal(&options)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", configs.TestingBasePath+"/meal/findByTag", nil)

	server.Handler.ServeHTTP(w, r)
	defer server.Close()

	defaultMeals := guessMealIDs(dao.GetDefaultFillMeals(
		guessTagIDs(dao.GetDefaultFillTags()),
		guessIngredientIDs(dao.GetDefaultFillIngredients()),
	))
	body := w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	// fmt.Println(string(body))
	meals := []api.Meal{}
	err := json.Unmarshal(body, &meals)
	s.NoError(err, "Body")
	s.ElementsMatch(defaultMeals, meals)
}

func (s *InfoTestSuite) TestMealFindMealsByTags() {
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     fmt.Sprintf("file:memdb%s%d?mode=memory&cache=shared&_pragma=foreign_keys(1)", regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(s.T().Name(), "_"), os.Getpid()),
		DbSample:  true,
		DbDebug:   false,
		JwtKey:    "5678",
	}
	server := s.buildServerMeal(&options)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", configs.TestingBasePath+"/meal/findByTag?tag=1", nil)

	server.Handler.ServeHTTP(w, r)
	defer server.Close()

	body := w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	//fmt.Println(string(body))
	meals := []api.Meal{}
	err := json.Unmarshal(body, &meals)
	s.NoError(err, "Body")
}

func (s *InfoTestSuite) TestMealDeleteMeal() {
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     fmt.Sprintf("file:memdb%s%d?mode=memory&cache=shared&_pragma=foreign_keys(1)", regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(s.T().Name(), "_"), os.Getpid()),
		DbSample:  true,
		DbDebug:   false,
		JwtKey:    "5678",
	}
	server := s.buildServerMeal(&options)
	token := s.getJwtToken(server)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", configs.TestingBasePath+"/meal/1", nil)
	r.Header.Add("Authorization", "Bearer "+token)

	server.Handler.ServeHTTP(w, r)
	defer server.Close()

	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+w.Body.String())
	//body := w.Body.Bytes()
	//fmt.Println(string(body))
}

func (s *InfoTestSuite) TestMealUpdateMeal() {
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     fmt.Sprintf("file:memdb%s%d?mode=memory&cache=shared&_pragma=foreign_keys(1)", regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(s.T().Name(), "_"), os.Getpid()),
		DbSample:  true,
		DbDebug:   false,
		JwtKey:    "5678",
	}
	server := s.buildServerMeal(&options)
	token := s.getJwtToken(server)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", configs.TestingBasePath+"/meal/1", nil)

	server.Handler.ServeHTTP(w, r)
	defer server.Close()

	body := w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	//fmt.Println(string(body))
	meal := &api.Meal{}
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

	server.Handler.ServeHTTP(w, r)
	defer server.Close()

	body = w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	//fmt.Println(string(body))
	mealUpdated := &api.Meal{}
	err = json.Unmarshal(body, mealUpdated)
	s.NoError(err, "Body")
	s.Equal("TEST", meal.Description, "Update")

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", configs.TestingBasePath+"/meal/1", nil)

	server.Handler.ServeHTTP(w, r)
	defer server.Close()

	body = w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	//fmt.Println(string(body))
	mealLatest := &api.Meal{}
	err = json.Unmarshal(body, mealLatest)
	s.NoError(err, "Body")
	s.Equal("TEST", mealLatest.Description, "Update Description")
	s.ElementsMatch(meal.Tags, mealLatest.Tags, "Update Tags")
	s.ElementsMatch(meal.Ingredients, mealLatest.Ingredients, "Update Ingredients")
}

func (s *InfoTestSuite) TestMealCreateMeal() {
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     fmt.Sprintf("file:memdb%s%d?mode=memory&cache=shared&_pragma=foreign_keys(1)", regexp.MustCompile("[^a-zA-Z0-9]+").ReplaceAllString(s.T().Name(), "_"), os.Getpid()),
		DbSample:  true,
		DbDebug:   false,
		JwtKey:    "5678",
	}
	server := s.buildServerMeal(&options)
	token := s.getJwtToken(server)
	tags := guessTagIDs(dao.GetDefaultFillTags())
	ingredients := guessIngredientIDs(dao.GetDefaultFillIngredients())
	meal := &api.Meal{
		Name:        "Vegan",
		Description: "Vegan pizza",
		PictureUrl:  "http://a.com",
		Price:       4.10,
		Kcal:        234,
		Ingredients: []api.Ingredient{
			ingredients[0],
			ingredients[5],
			ingredients[6],
		},
		Tags: []api.Tag{
			tags[2],
		},
	}
	body, err := json.Marshal(meal)
	s.NoError(err, "Body")
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", configs.TestingBasePath+"/meal/0", bytes.NewReader(body))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+token)

	server.Handler.ServeHTTP(w, r)
	defer server.Close()

	body = w.Body.Bytes()
	s.Equal(http.StatusOK, w.Result().StatusCode, "Status, "+string(body))
	//fmt.Println(string(body))
	mealCreated := &api.Meal{}
	err = json.Unmarshal(body, mealCreated)
	s.NoError(err, "Body")
}

func guessTagIDs(tags []api.Tag) []api.Tag {
	for t := range tags {
		tags[t].Id = int64(t + 1)
	}

	return tags
}

func guessIngredientIDs(ingredients []api.Ingredient) []api.Ingredient {
	for i := range ingredients {
		ingredients[i].Id = int64(i + 1)
	}

	return ingredients
}

func guessMealIDs(meals []api.Meal) []api.Meal {
	for m := range meals {
		meals[m].Id = int64(m + 1)
	}

	return meals
}
