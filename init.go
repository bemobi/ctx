package webctx

import (
	"net/http"

	"github.com/auth0/go-jwt-middleware"
	"github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func Init(router *mux.Router, privKey, pubKey string, vars map[string]interface{}) error {
	LoadSecureKeys(privKey, pubKey)

	// create and configure a new context
	ctx := &Context{
		Router: router,
		Vars:   vars,

		middleware: jwtmiddleware.New(
			jwtmiddleware.Options{
				ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) { return verifyKey, nil },
			},
		),
	}

	// add the registered endpoints
	for _, endpoint := range endpoints {
		httpHandler := newContextHandler(ctx, endpoint)
		if endpoint.Public {
			ctx.Router.HandleFunc(endpoint.Path, httpHandler)
		} else {
			ctx.Router.Handle(
				endpoint.Path, negroni.New(
					negroni.HandlerFunc(ctx.middleware.HandlerWithNext),
					negroni.Wrap(http.HandlerFunc(httpHandler)),
				),
			)
		}
	}

	return nil
}
