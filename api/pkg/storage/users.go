package storage

import (
	"database/sql"
	"fmt"

  types "api/pkg/types"
)

func (s *PostgresStore) CreateUser(user *types.User) error {
	query := `INSERT INTO users 
	(email, username, encrypted_password)
	VALUES ($1, $2, $3);`

	_, err := s.db.Query(
		query,
		user.Email,
		user.Username,
		user.EncryptedPassword)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) UpdateUser(*types.User) error {
	return nil
}

func (s *PostgresStore) DeleteUser(id string) error {
	_, err := s.db.Query("DELETE FROM users WHERE id = $1", id)
	return err
}

func scanIntoUser(rows *sql.Rows) (*types.User, error) {
	user:= new(types.User)
	err := rows.Scan(
    &user.ID,
		&user.Email,
		&user.Username,
		&user.EncryptedPassword,
    &user.CreatedAt)

	return user, err
}

func (s *PostgresStore) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("user with email [%s] not found", email)
}

func (s *PostgresStore) GetUserByID(id string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("user %d not found", id)
}

func (s *PostgresStore) GetUserByUsername(username string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoUser(rows)
	}

	return nil, fmt.Errorf("user %d not found", username)
}
