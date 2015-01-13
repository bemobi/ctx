package ctx

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/nicksnyder/go-i18n/i18n"
)

type Context struct {
	Router *mux.Router
	Vars   map[string]interface{}

	T i18n.TranslateFunc

	Token      *jwt.Token
	middleware *jwtmiddleware.JWTMiddleware
}
