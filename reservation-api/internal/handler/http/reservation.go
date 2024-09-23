package http

import (
	reservation "reservation-api/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *reservation.Service
}

func New(s *reservation.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Get(c *gin.Context) { // return
	h.svc.Get(c)
}

func (h *Handler) Create(c *gin.Context) { // return
	h.svc.Create(c)
}

func (h *Handler) Delete(c *gin.Context) { // return
	h.svc.Delete(c)
}
