package app

import (
	"errors"
)

var (
	// ErrNotFound is an implementation agnostic error that should be returned
	// by any service implementation when a record was not located.
	ErrNotFound = errors.New("app: resource requested could not be found")
)

// UserService defines the interface used to interact with the users datastore.
// Implementations can be found in packages like the psql package.
type UserService interface {
	ByEmail(email string) (*User, error)
	ByToken(token string) (*User, error)
	UpdateToken(userID int, token string) error
}

// WidgetService defines the interface used to interact with the widget
// datastore. Implementations can be found in packages like the psql package.
type WidgetService interface {
	ByUser(userID int) ([]Widget, error)
	Create(widget *Widget) error
}
