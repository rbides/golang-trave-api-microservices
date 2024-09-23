package http

import (
	"log"
	"net/http"
	user "users-api/internal/service"
	request "users-api/pkg/models/requests"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	svc *user.Service
}

func New(s *user.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Get(c *gin.Context) {
	log.Println("Get Handler")
	params := request.QueryParams{}
	if err := c.ShouldBindQuery(&params); err != nil {
		log.Println("Validation Error: ", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	log.Println(params)
	users, err := h.svc.Get(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetById(c *gin.Context) {
	log.Println("GetById Handler")
	id, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		log.Println("Invalid ID", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	u, err := h.svc.GetById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, u)
}

func (h *Handler) Create(c *gin.Context) {
	log.Println("Create Handler")
	r := request.CreateUserRequest{}

	if err := c.ShouldBindJSON(&r); err != nil {
		log.Println("Validation Error: ", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := h.svc.Create(c, &r); err != nil {
		c.JSON(http.StatusInternalServerError, "Failed creating user.")
		return
	}
	c.JSON(http.StatusOK, r)
}

func (h *Handler) Update(c *gin.Context) {
	log.Println("Update Handler")
	h.svc.Update(c)
}

func (h *Handler) Delete(c *gin.Context) {
	log.Println("Del Handler")
	id, err := uuid.Parse(c.Param("userID"))
	if err != nil {
		log.Println("Invalid ID", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err = h.svc.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, "Deletion failed.")
		return
	}
	c.JSON(http.StatusOK, "User deleted.")
}
