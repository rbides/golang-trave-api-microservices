package travel

import (
	"log"
	"time"
	travel "travel-api/internal/repository/postgresql"
	request "travel-api/pkg/models/requests"
	models "travel-api/pkg/models/travel"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service struct {
	repo *travel.Repository
}

func New(r *travel.Repository) *Service {
	return &Service{r}
}

// TODO: get seats
// TODO: Change seats to processing

func (s *Service) Get(c *gin.Context) ([]models.Travel, error) { // return
	log.Println("Get Service")
	travels, err := s.repo.Get(c)
	if err != nil {
		log.Println("error", err)
		return nil, err
	}
	return travels, nil

}

func (s *Service) GetById(c *gin.Context, id uuid.UUID) (models.Travel, error) { // return
	log.Println("GetById Service")
	t, err := s.repo.GetById(c, id)
	if err != nil {
		log.Println("error", err)
		return models.Travel{}, err
	}

	return t, nil
}

func (s *Service) Create(c *gin.Context, t *request.TravelRequest) error {
	log.Println("Create Service", time.Now())
	travel := models.Travel{
		ID:          uuid.New(),
		Name:        t.Name,
		Destination: t.Destination,
		Price:       models.Money(t.Price),
		Departure:   t.Departure,
	}
	seats := make([]models.Seat, t.Seats)
	for i := range t.Seats {
		seats[i].ID = uuid.New()
		seats[i].TravelID = travel.ID
		seats[i].Position = i + 1
	}

	return s.repo.Create(c, travel, seats)
}

func (s *Service) Update(c *gin.Context) {
	log.Println("Update Service")
	s.repo.Update()
}

func (s *Service) Delete(c *gin.Context, id uuid.UUID) error {
	log.Println("Del Service")
	return s.repo.Delete(c, id)
}

func (s *Service) GetSeats(c *gin.Context, travel_id uuid.UUID) ([]models.Seat, error) {
	log.Println("Get Seats Service")
	seats, err := s.repo.GetSeats(c, travel_id)
	if err != nil {
		return nil, err
	}
	return seats, nil
}

func (s *Service) UpdateSeats(c *gin.Context, r request.UpdateSeatsRequest) error {
	log.Println("Update Seats Service")
	if err := s.repo.UpdateSeats(c, r); err != nil {
		return err
	}

	return nil
}
