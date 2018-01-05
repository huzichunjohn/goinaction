package handlers

import (
	"net/http"

	"github.com/takama/bit"
)

func (h *Handler) Ready(c bit.Control) {
	c.Code(http.StatusOK)
	c.Body(http.StatusText(http.StatusOK))
}
