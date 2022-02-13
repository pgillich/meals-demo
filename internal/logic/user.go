package logic

import (
	"fmt"
	"log"
	"strings"
	"time"

	"emperror.dev/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"

	"github.com/pgillich/meals-demo/configs"
	"github.com/pgillich/meals-demo/internal/dao"
	"github.com/pgillich/meals-demo/internal/models"
	"github.com/pgillich/meals-demo/internal/restapi/operations"
	"github.com/pgillich/meals-demo/internal/restapi/operations/user"
)

func SetUserAPI(config configs.Options, api *operations.OpenAPIFoodstoreAPI) {
	userAPI := &UserAPI{
		jwtKey: config.JwtKey,
		method: jwt.SigningMethodHS256,
	}
	var err error

	userAPI.dbHandler, err = dao.NewHandler(config)
	if err != nil {
		log.Fatal(err)
	}

	api.JWTAuth = userAPI.ValidateHeader
	api.APIKeyAuthenticator = security.APIKeyAuth
	api.UserLoginHandler = user.LoginHandlerFunc(userAPI.Login)
}

type UserAPI struct {
	dbHandler *dao.Handler
	jwtKey    string
	method    jwt.SigningMethod
}

func (userAPI *UserAPI) Login(params user.LoginParams) middleware.Responder {
	if params.Login.Email == nil || params.Login.Password == nil {
		return user.NewLoginInternalServerError().WithPayload("Missing cred info")
	}
	authUser, err := userAPI.dbHandler.AuthenticateUser(*params.Login.Email, *params.Login.Password)
	if err != nil {
		return user.NewLoginNotFound()
	}

	token, err := userAPI.GenerateJWT(authUser.Email)
	if err != nil {
		return user.NewLoginInternalServerError().WithPayload("Error defining token")
	}
	return user.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: token})
}

func (userAPI *UserAPI) ValidateHeader(bearerHeader string) (*models.User, error) {
	bearerToken := strings.Split(bearerHeader, " ")[1]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(userAPI.jwtKey), nil
	})
	if err != nil {
		return nil, errors.WrapWithDetails(err, "Unable to parse JWT")
	}
	if token.Valid {
		return userAPI.dbHandler.GetUserByEmail(fmt.Sprintf("%v", claims["email"]))
	}
	return nil, errors.New("invalid token")
}

func (userAPI *UserAPI) GenerateJWT(email string) (string, error) {
	tokenString, err := jwt.NewWithClaims(userAPI.method, jwt.MapClaims{
		"authorized": true,
		"email":      email,
		"exp":        time.Now().Add(time.Minute * 300).Unix(),
	}).SignedString([]byte(userAPI.jwtKey))
	if err != nil {
		return "", errors.WrapWithDetails(err, "Error generating Token")
	}

	return tokenString, nil
}
