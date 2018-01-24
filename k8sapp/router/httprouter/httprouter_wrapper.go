package httprouter

import (
	"net/http"

	"goinaction/k8sapp/router"

	"github.com/julienschmidt/httprouter"
)

type httpRouter struct {
	httprouter.Router
}

func New() router.HTTPRouter {
	router := new(httpRouter)
	router.RedirectTrailingSlash = true
	router.RedirectFixedPath = true
	router.HandleMethodNotAllowed = true
	router.HandleOPTIONS = true
	return router
}

func (hr *httpRouter) UseOptionsReplies(enabled bool) {
	hr.HandleOPTIONS = enabled
}

func (hr *httpRouter) SetupNotAllowedHandler(h http.Handler) {
	hr.MethodNotAllowed = h
}

func (hr *httpRouter) SetupNotFoundHandler(h http.Handler) {
	hr.NotFound = h
}

func (hr *httpRouter) SetupRecoveryHandler(f func(http.ResponseWriter, *http.Request, interface{})) {
	hr.PanicHandler = f
}

func (hr *httpRouter) Listen(hostPort string) error {
	return http.ListenAndServe(hostPort, hr)
}
