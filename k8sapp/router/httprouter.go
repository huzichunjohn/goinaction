package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HTTPRouter interface {
	GET(path string, h httprouter.Handle)
	PUT(path string, h httprouter.Handle)
	POST(path string, h httprouter.Handle)
	DELETE(path string, h httprouter.Handle)
	HEAD(path string, h httprouter.Handle)
	OPTIONS(path string, h httprouter.Handle)
	PATCH(path string, h httprouter.Handle)

	UseOptionsReplies(bool)

	SetupNotAllowedHandler(http.Handler)

	SetupNotFoundHandler(http.Handler)

	SetupRecoveryHandler(func(http.ResponseWriter, *http.Request, interface{}))

	Listen(hostPort string) error
}
