package travel

import (
	"log"
	"net/http"
	"time"
	travel "travel-api/internal/service"
	request "travel-api/pkg/models/requests"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	svc *travel.Service
}

func New(s *travel.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Get(c *gin.Context) {
	log.Println("Get Handler")
	travels, err := h.svc.Get(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, travels)
}

func (h *Handler) GetById(c *gin.Context) {
	//TODO: add 404 for not found
	log.Println("GetById Handler")
	id, err := uuid.Parse(c.Param("travelID"))
	if err != nil {
		log.Println("Validation Error: ", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	t, err := h.svc.GetById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, t)
}

func (h *Handler) Create(c *gin.Context) {
	log.Println("Create Handler")
	t := request.TravelRequest{}

	if err := c.ShouldBindJSON(&t); err != nil {
		log.Println("Error json", err)
		c.JSON(http.StatusBadRequest, "Invalid input data")
		return
	}
	// Checks if departure date is > Now
	if b := t.Departure.After(time.Now()); !b {
		log.Println("Invalid departure date", t.Departure) // create a custom error for this
		c.JSON(http.StatusBadRequest, "Invalid departure date")
		return
	}

	if err := h.svc.Create(c, &t); err != nil {
		c.JSON(http.StatusInternalServerError, "Failed creating Travel.")
		return
	}
	c.JSON(http.StatusOK, "Creation success")
}

func (h *Handler) Update(c *gin.Context) {
	log.Println("Update Handler")
	h.svc.Update(c)
}

func (h *Handler) Delete(c *gin.Context) {
	log.Println("Del Handler")
	id, err := uuid.Parse(c.Param("travelID"))
	if err != nil {
		log.Println("Invalid ID", err)
		return
	}
	if err = h.svc.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, "Failed deleting travel")
		return
	}

	c.JSON(http.StatusOK, "Deletion successfull")
}

func (h *Handler) GetSeats(c *gin.Context) {
	log.Println("Get Handler")
	travel_id, err := uuid.Parse(c.Param("travelID"))
	if err != nil {
		log.Println("Validation Error: ", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	travels, err := h.svc.GetSeats(c, travel_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, travels)
}

func (h *Handler) UpdateSeats(c *gin.Context) {
	log.Println("Get Handler")
	r := request.UpdateSeatsRequest{}
	if err := c.ShouldBindBodyWithJSON(&r); err != nil {
		log.Println("Validation Error: ", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err := h.svc.UpdateSeats(c, r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "Seats updated")
}
