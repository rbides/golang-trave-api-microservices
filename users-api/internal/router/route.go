package router

import (
	"users-api/internal/handler/http"
	"users-api/internal/repository/postgresql"
	user "users-api/internal/service"

	"github.com/gin-gonic/gin"
)

func initUsers(r *gin.RouterGroup) {
	repo := postgresql.New()
	svc := user.New(repo)
	handler := http.New(svc)
	users := r.Group("/users")
	users.GET("", handler.Get)
	users.POST("", handler.Create)
	users.GET(":userID", handler.GetById)
	users.PUT(":userID", handler.Update)
	users.DELETE(":userID", handler.Delete)
}

func Init() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	initUsers(api)
	return r
}
