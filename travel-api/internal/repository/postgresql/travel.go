package travel

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"travel-api/internal/repository"
	request "travel-api/pkg/models/requests"
	models "travel-api/pkg/models/travel"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func New() *Repository {
	godotenv.Load()
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Panic(err)
	}

	return &Repository{db}
}

func (r *Repository) Get(c *gin.Context) ([]models.Travel, error) {
	log.Println("Repo - Getting travels")
	rows, err := r.db.QueryContext(c, "SELECT * FROM travels")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var res []models.Travel
	for rows.Next() {
		t := models.Travel{}
		if err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.Destination,
			&t.Price,
			&t.Departure,
			&t.CreatedAt,
			&t.UpdatedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				log.Println("Err Not found")
				return nil, repository.ErrNotFound
			}
			return nil, err
		}
		res = append(res, t)
	}
	if err := rows.Err(); err != nil {
		log.Println("Query error: ", err)
		return nil, err
	}
	return res, nil
}

func (r *Repository) GetById(c *gin.Context, id uuid.UUID) (models.Travel, error) {
	log.Println("Repo - Getting travel: ", id)
	row := r.db.QueryRowContext(c, "SELECT * FROM travels WHERE id=$1", id)
	var t models.Travel
	if err := row.Scan(
		&t.ID,
		&t.Name,
		&t.Destination,
		&t.Price,
		&t.Departure,
		&t.CreatedAt,
		&t.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Println("Err Not found")
			return t, repository.ErrNotFound
		}
		return t, err
	}
	return t, nil
}

// Create a Travel and its corresponding Seats
func (r *Repository) Create(c *gin.Context, t models.Travel, seats []models.Seat) error {
	log.Println("Creating travel with id: ", t)

	// TODO: handle errors

	tx, err := r.db.BeginTx(c, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(
		c,
		"INSERT INTO travels (id, name, destination, price, departure) VALUES ($1, $2, $3, $4, $5);",
		t.ID,
		t.Name,
		t.Destination,
		t.Price,
		t.Departure,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	// User CopyIn for faster bulk inserting seats
	stmt, err := tx.PrepareContext(c, pq.CopyIn("seats", "id", "travel_id", "position"))
	if err != nil {
		log.Println(err)
		return err
	}
	for _, s := range seats {
		_, err = stmt.ExecContext(
			c,
			s.ID,
			s.TravelID,
			s.Position,
		)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	_, err = stmt.ExecContext(c)
	if err != nil {
		log.Println(err)
		return err
	}

	err = stmt.Close()
	if err != nil {
		log.Println(err)
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Println(err) // create error
		return err
	}

	return nil
}

// TODO
func (r *Repository) Update() {
	log.Println("Update")
	// create error
}

// Delete a travel and it's corresponding seats
func (r *Repository) Delete(c *gin.Context, id uuid.UUID) error {
	log.Println("Deleting travel with id: ", id, " and it's corresponding seats")
	tx, err := r.db.BeginTx(c, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(
		c,
		"DELETE FROM seats WHERE travel_id = $1",
		id,
	)
	if err != nil {
		log.Println(err) // create error
		return err
	}
	_, err = tx.ExecContext(
		c,
		"DELETE FROM travels WHERE id = $1",
		id,
	)
	if err != nil {
		log.Println(err) // create error
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Println(err) // create error
		return err
	}

	return nil
}

func (r *Repository) GetSeats(c *gin.Context, travel_id uuid.UUID) ([]models.Seat, error) {
	log.Println("Repo - Get seats")
	rows, err := r.db.QueryContext(c, "SELECT * FROM seats WHERE travel_id = $1", travel_id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var res []models.Seat
	for rows.Next() {
		s := models.Seat{}
		if err := rows.Scan(
			&s.ID,
			&s.TravelID,
			&s.Position,
			&s.Status,
			&s.CreatedAt,
			&s.UpdatedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				log.Println("Err Not found")
				return nil, repository.ErrNotFound
			}
			return nil, err
		}
		res = append(res, s)
	}
	if err := rows.Err(); err != nil {
		log.Println("Query error: ", err)
		return nil, err
	}
	return res, nil
}

func (r *Repository) UpdateSeats(c *gin.Context, req request.UpdateSeatsRequest) error {
	log.Println("Update Seats repo")
	var ids []string
	for _, s := range req.Seats {
		ids = append(ids, fmt.Sprintf("'%v'", s.String()))
	}
	query := fmt.Sprintf(`UPDATE seats SET status=$1 WHERE id IN (%v)`, strings.Join(ids, ","))
	log.Println(query)

	_, err := r.db.ExecContext(
		c,
		query,
		req.Status,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
