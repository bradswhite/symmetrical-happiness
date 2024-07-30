package storage

import (
	"database/sql"
  "os"
	"fmt"
  "log"

  _ "github.com/lib/pq"
  types "api/pkg/types"
)

type Storage interface {
	CreateUser(*types.User) error
	GetUserByID(string) (*types.User, error)
	GetUserByEmail(string) (*types.User, error)
	UpdateUser(*types.User) error
	DeleteUser(string) error

  CreateSoftware(*types.Software) error
  GetSoftware() ([]*types.Software, error)
  GetSoftwareByID(string) (*types.Software, error)
	UpdateSoftware(*types.Software) error
	DeleteSoftware(string) error

  CreateSoftwareLike(*types.SoftwareLike) error
  GetSoftwareLikesBySoftware(string) ([]*types.SoftwareLike, error)
	DeleteSoftwareLike(string, string) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
  user := os.Getenv("DB_USER")
  dbname := os.Getenv("DB_NAME")
  host := os.Getenv("DB_HOST")
  port := os.Getenv("DB_PORT")
  password := os.Getenv("DB_PASS")

  connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createTables()
}

func (s *PostgresStore) createTables() error {
  content, err := os.ReadFile("./schema.sql")
  if err != nil {
      log.Fatal(err)
      return err
  }

	_, err = s.db.Exec(string(content))
  if err != nil {
    log.Fatal(err)
  }
	return err
}
