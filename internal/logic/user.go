package logic

import (
	"context"
	"strings"
	"time"

	"emperror.dev/errors"
	"github.com/dgrijalva/jwt-go"

	"github.com/pgillich/meals-demo/internal/api"
)

func (fs *FoodStore) Login(ctx context.Context, request api.LoginRequestObject) (api.LoginResponseObject, error) {
	if request.Body == nil || request.Body.Email == "" || request.Body.Password == "" {
		return api.Login400JSONResponse{Message: "Missing cred info"}, nil
	}
	authUser, err := fs.dbHandler.AuthenticateUser(request.Body.Email, request.Body.Password)
	if err != nil {
		return api.Login404AsteriskResponse{}, nil //nolint:nilerr // in another field
	}

	token, err := fs.GenerateJWT(authUser.Email)
	if err != nil {
		//nolint:nilerr // in another field
		return api.Login500AsteriskResponse{Body: strings.NewReader("Error defining token")}, nil
	}

	return api.Login200JSONResponse{Success: true, Token: token}, nil
}

func (fs *FoodStore) ValidateHeader(bearerHeader string) (*api.User, error) {
	bearerToken := strings.Split(bearerHeader, " ")[1]
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(fs.jwtKey), nil
	})
	if err != nil {
		return nil, errors.WrapWithDetails(err, "Unable to parse JWT")
	}
	if token.Valid {
		return fs.dbHandler.GetUserByEmail(claims.Subject)
	}

	return nil, errors.New("invalid token")
}

func (fs *FoodStore) GenerateJWT(email string) (string, error) {
	tokenString, err := jwt.NewWithClaims(fs.method, jwt.StandardClaims{
		Subject:   email,
		ExpiresAt: time.Now().Add(fs.jwtExpireSec).Unix(),
		Issuer:    "jwtIssuer",
	}).SignedString([]byte(fs.jwtKey))
	if err != nil {
		return "", errors.WrapWithDetails(err, "Error generating Token")
	}

	return tokenString, nil
}
