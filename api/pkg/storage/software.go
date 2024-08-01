package storage

import (
	"database/sql"
	"fmt"

  //pq "github.com/lib/pq"
  types "api/pkg/types"
)

func (s *PostgresStore) CreateSoftware(software *types.Software) error {
	query := `INSERT INTO software
	(name, title, description, image, url, username)
	VALUES ($1, $2, $3, $4, $5, $6);`

	_, err := s.db.Query(
		query,
		software.Name,
		software.Title,
		software.Description,
		software.Image,
		software.Url,
		//pq.Array(Software.Alts),
		software.Username)

	if err != nil {
		return err
	}

	return nil
}


func (s *PostgresStore) GetSoftwareByID(id string) (*types.Software, error) {
	rows, err := s.db.Query("SELECT * FROM software WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoSoftware(rows)
	}

	return nil, fmt.Errorf("software %d not found", id)
}

func (s *PostgresStore) GetSoftware() ([]*types.Software, error) {
	rows, err := s.db.Query("SELECT * FROM software")
	if err != nil {
		return nil, err
	}

	softwares := []*types.Software{}
	for rows.Next() {
		software, err := scanIntoSoftware(rows)
		if err != nil {
			return nil, err
		}
		softwares = append(softwares, software)
	}

	return softwares, nil
}

func scanIntoSoftware(rows *sql.Rows) (*types.Software, error) {
	software := new(types.Software)
	err := rows.Scan(
    &software.ID,
		&software.Name,
		&software.Title,
		&software.Description,
		&software.Image,
		&software.Url,
		&software.Username,
		&software.CreatedAt)

	return software, err
}

func (s *PostgresStore) UpdateSoftware(softwareId string, software *types.UpdateSoftwareRequest) error {
	query := `UPDATE software
    SET name=$2, title=$3, description=$4, image=$5, url=$6
    WHERE id=$1;`

	_, err := s.db.Exec(
		query,
    softwareId,
		software.Name,
		software.Title,
		software.Description,
		software.Image,
		software.Url)
	if err != nil {
		return err
	}
 
	return nil
}

func (s *PostgresStore) DeleteSoftware(id string) error {
	_, err := s.db.Query("DELETE FROM software WHERE id = $1", id)
	return err
}
