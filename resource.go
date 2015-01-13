package webctx

import "net/http"

type Getter interface {
	GET(c *Context, rw http.ResponseWriter, req *http.Request) error
}

type Poster interface {
	POST(c *Context, rw http.ResponseWriter, req *http.Request) error
}

type Putter interface {
	PUT(c *Context, rw http.ResponseWriter, req *http.Request) error
}

type Deleter interface {
	DELETE(c *Context, rw http.ResponseWriter, req *http.Request) error
}

func Resource(path string, res interface{}, public bool) {
	handlers := make(map[string]ContextHandler, 0)

	if handler, implements := res.(Getter); implements {
		handlers["GET"] = handler.GET
	}

	if handler, implements := res.(Poster); implements {
		handlers["POST"] = handler.POST
	}

	if handler, implements := res.(Putter); implements {
		handlers["PUT"] = handler.PUT
	}

	if handler, implements := res.(Deleter); implements {
		handlers["DELETE"] = handler.DELETE
	}

	Register(&Endpoint{Path: path, Public: public, Handlers: handlers})
}

func (e *Endpoint) GET(handler ContextHandler) *Endpoint { e.Handlers["GET"] = handler; return e }

func (e *Endpoint) POST(handler ContextHandler) *Endpoint { e.Handlers["POST"] = handler; return e }

func (e *Endpoint) PUT(handler ContextHandler) *Endpoint { e.Handlers["PUT"] = handler; return e }

func (e *Endpoint) DELETE(handler ContextHandler) *Endpoint { e.Handlers["DELETE"] = handler; return e }
