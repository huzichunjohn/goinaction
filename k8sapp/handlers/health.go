package handlers

import (
	"net/http"

	"github.com/takama/bit"
)

func (h *Handler) Health(c bit.Control) {
	c.Code(http.StatusOK)
	c.Body(http.StatusText(http.StatusOK))
}
