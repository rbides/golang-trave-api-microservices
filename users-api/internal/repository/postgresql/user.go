package postgresql

import (
	"database/sql"
	"log"
	"os"
	"users-api/internal/repository"
	request "users-api/pkg/models/requests"
	models "users-api/pkg/models/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

func (r *Repository) Get(c *gin.Context, params request.QueryParams) ([]models.User, error) {
	log.Println("Repo - Getting users")
	query := "SELECT * FROM users"
	args := []interface{}{}
	if params.Email != "" {
		query += " WHERE email = $1"
		args = append(args, params.Email)
	}
	rows, err := r.db.QueryContext(c, query, args...)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var res []models.User
	for rows.Next() {
		u := models.User{}
		if err := rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.CreatedAt,
			&u.UpdatedAt,
		); err != nil {
			if err == sql.ErrNoRows {
				log.Println("Err Not found")
				return nil, repository.ErrNotFound
			}
			return nil, err
		}
		res = append(res, u)
	}
	if err := rows.Err(); err != nil {
		log.Println("Query error: ", err)
		return nil, err
	}
	return res, nil
}

func (r *Repository) GetById(c *gin.Context, id uuid.UUID) (models.User, error) {
	log.Println("Repo - Getting user: ", id)
	row := r.db.QueryRowContext(
		c,
		"SELECT * FROM users WHERE id=$1",
		id,
	)
	var u models.User
	if err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.CreatedAt,
		&u.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Println("Err Not found")
			return u, repository.ErrNotFound
		}
		return u, err
	}
	return u, nil
}

// Create User
func (r *Repository) Create(c *gin.Context, u models.User) error {
	log.Println("Creating user with id: ", u.ID)

	// TODO: handle errors

	_, err := r.db.ExecContext(
		c,
		"INSERT INTO users (id, name, email) VALUES ($1, $2, $3);",
		u.ID,
		u.Name,
		u.Email,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *Repository) Update() {
	// TODO
	log.Println("Update")
	// create error
}

func (r *Repository) Delete(c *gin.Context, id uuid.UUID) error {
	log.Println("Del")
	res, err := r.db.ExecContext(
		c,
		"DELETE FROM users WHERE id = $1",
		id,
	)
	if err != nil {
		log.Println(err) // create error
		return err
	}
	if n, err := res.RowsAffected(); n == 0 || err != nil {
		log.Println(repository.ErrNotFound)
		return repository.ErrNotFound
	}

	return nil
}
