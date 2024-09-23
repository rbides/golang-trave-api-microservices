package reservation

import (
	"reservation-api/internal/repository/postgresql"

	"github.com/gin-gonic/gin"
)

type Service struct {
	repo *postgresql.Repository
}

func New(r *postgresql.Repository) *Service {
	return &Service{r}
}

func (s *Service) Get(c *gin.Context) { // return
	s.repo.Get(c)
}

func (s *Service) Create(c *gin.Context) { // return
	s.repo.Create(c)
}

func (s *Service) Delete(c *gin.Context) { // return
	s.repo.Delete(c)
}
