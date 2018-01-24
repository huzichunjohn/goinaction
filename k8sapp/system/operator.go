package system

import (
	"errors"
)

var ErrNotImplemented = errors.New("This method is not implemented")

type Operator interface {
	Reload() error
	Maintenance() error
	Shutdown() error
}

type Handling struct{}

func (h Handling) Reload() error {
	return ErrNotImplemented
}

func (h Handling) Maintenance() error {
	return ErrNotImplemented
}

func (h Handling) Shutdown() error {
	return ErrNotImplemented
}
