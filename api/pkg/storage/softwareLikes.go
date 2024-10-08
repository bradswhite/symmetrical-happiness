package storage

import (
	"database/sql"
	"fmt"

  types "api/pkg/types"
)

func (s *PostgresStore) CreateSoftwareLike(softwareLike *types.SoftwareLike) error {
	query := `INSERT INTO software_likes 
	(software_id, username)
	VALUES ($1, $2);`

	_, err := s.db.Query(
		query,
    softwareLike.SoftwareID,
    softwareLike.Username)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetSoftwareLikesBySoftware(softwareId string) ([]*types.SoftwareLike, error) {
	rows, err := s.db.Query("SELECT * from software_likes WHERE software_id = $1;", softwareId)
	if err != nil {
    fmt.Println(err)
		return nil, err
	}

	softwareLikes := []*types.SoftwareLike{}
	for rows.Next() {
		softwareLike, err := scanIntoSoftwareLikes(rows)
		if err != nil {
			return nil, err
		}
		softwareLikes = append(softwareLikes, softwareLike)
	}

	return softwareLikes, nil
}

func scanIntoSoftwareLikes(rows *sql.Rows) (*types.SoftwareLike, error) {
	softwareLike := new(types.SoftwareLike)
	err := rows.Scan(
		&softwareLike.SoftwareID,
		&softwareLike.Username,
		&softwareLike.LikedAt)

	return softwareLike, err
}

func (s *PostgresStore) DeleteSoftwareLike(softwareId, username string) error {
	_, err := s.db.Query("DELETE FROM software_likes WHERE username = $1 and softwareID = $2", username, softwareId)
	return err
}
