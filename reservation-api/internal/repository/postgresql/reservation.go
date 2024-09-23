package postgresql

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
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

func (r *Repository) Get(c *gin.Context) {
	// Add filter by user and travel
}

func (r *Repository) Create(c *gin.Context) {
}

func (r *Repository) Delete(c *gin.Context) {
	// TODO: delete by user + travel id
}
