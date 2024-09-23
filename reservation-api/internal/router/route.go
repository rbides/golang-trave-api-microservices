package router

import (
	"reservation-api/internal/handler/http"
	"reservation-api/internal/repository/postgresql"
	reservation "reservation-api/internal/service"

	"github.com/gin-gonic/gin"
)

func initReservations(r *gin.RouterGroup) {
	repo := postgresql.New()
	svc := reservation.New(repo)
	handler := http.New(svc)
	reservations := r.Group("/reservation")
	reservations.GET("", handler.Get)
	reservations.POST("", handler.Create)
	reservations.DELETE(":travelID/:userID", handler.Delete)
}

func Init() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	initReservations(api)
	return r
}
