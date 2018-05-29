package http

import (
	"net/http"
	app "widgetapp"

	"github.com/gorilla/mux"
)

// NewServer will construct a Server and apply all of the necessary routes
func NewServer(us app.UserService, ws app.WidgetService) *Server {
	server := Server{
		authMw: &htmlAuthMw{
			userService: us,
		},
		users:   htmlUserHandler(us),
		widgets: htmlWidgetHandler(ws),
		router:  mux.NewRouter(),
	}
	server.routes()
	return &server
}

// Server is our HTTP server with routes for all our endpoints.
//
// The zero value is NOT useful - you should use the NewServer function
// to create a server.
type Server struct {
	// unexported types - do not use the zero value
	authMw  AuthMw
	users   *UserHandler
	widgets *WidgetHandler
	router  *mux.Router
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) routes() {
	s.router.Handle("/", http.RedirectHandler("/signin", http.StatusFound))

	// User routes
	s.router.HandleFunc("/signin", s.users.ShowSignin).Methods("GET")
	s.router.HandleFunc("/signin", s.users.ProcessSignin).Methods("POST")

	// Widget routes
	s.router.Handle("/widgets", ApplyMwFn(s.widgets.Index,
		s.authMw.SetUser, s.authMw.RequireUser)).Methods("GET")
	s.router.Handle("/widgets", ApplyMwFn(s.widgets.Create,
		s.authMw.SetUser, s.authMw.RequireUser)).Methods("POST")
	s.router.Handle("/widgets/new", ApplyMwFn(s.widgets.New,
		s.authMw.SetUser, s.authMw.RequireUser)).Methods("GET")
}
