package ctx

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/nicksnyder/go-i18n/i18n"
)

type ContextHandler func(c *Context, rw http.ResponseWriter, req *http.Request) error

func newContextHandler(context *Context, endpoint *Endpoint) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {

		// check if the http method is mapped
		var handler ContextHandler
		var found bool
		if handler, found = endpoint.Handlers[req.Method]; !found {
			http.Error(rw, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// extract the token, if the endpoint is private
		if !endpoint.Public {
			context.Token, _ = jwt.ParseFromRequest(
				req, context.middleware.Options.ValidationKeyGetter,
			)
		}

		// update T
		acceptLang := req.Header.Get("Accept-Language")
		defaultLang := "en-US"
		context.T = i18n.MustTfunc(acceptLang, defaultLang)

		// call the apropriate handler
		err := handler(context, rw, req)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
		}
	}
}
