package user

import (
	"log"
	"time"
	"users-api/internal/repository/postgresql"
	request "users-api/pkg/models/requests"
	models "users-api/pkg/models/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service struct {
	repo *postgresql.Repository
}

func New(r *postgresql.Repository) *Service {
	return &Service{r}
}

func (s *Service) Get(c *gin.Context, params request.QueryParams) ([]models.User, error) { // return
	log.Println("Get Service")
	users, err := s.repo.Get(c, params)
	if err != nil {
		log.Println("error", err)
	}
	return users, nil
}

func (s *Service) GetById(c *gin.Context, id uuid.UUID) (models.User, error) { // return
	log.Println("GetById Service")
	user, err := s.repo.GetById(c, id)
	if err != nil {
		log.Println("error", err)
		return models.User{}, err
	}
	return user, nil
}

func (s *Service) Create(c *gin.Context, u *request.CreateUserRequest) error {
	log.Println("Create Service", time.Now())
	user := models.User{
		ID:    uuid.New(),
		Name:  u.Name,
		Email: u.Email,
	}
	return s.repo.Create(c, user)
}

func (s *Service) Update(c *gin.Context) {
	log.Println("Update Service")
	s.repo.Update()
}

func (s *Service) Delete(c *gin.Context, id uuid.UUID) error {
	log.Println("Del Service")
	return s.repo.Delete(c, id)
}
