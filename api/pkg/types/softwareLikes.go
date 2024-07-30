package types

import (
	"time"

	"github.com/google/uuid"
)

type SoftwareLikeRequest struct {
	SoftwareID   string    `json:"softwareId"`
	Username     string    `json:"username"`
}

type CreateSoftwareLikeRequest struct {
	SoftwareID   string   `json:"softwareId"`
	Username     string   `json:"username"`
}

type SoftwareLike struct {
	SoftwareID   uuid.UUID   `json:"softwareId"`
	Username     string      `json:"username"`
	LikedAt      time.Time   `json:"likedAt"`
}

func NewSoftwareLike(softwareId, username string) (*SoftwareLike, error) {
  softwareUuid, err := uuid.Parse(softwareId)
  if err != nil {
	  return nil, err	
	}

	return &SoftwareLike{
    SoftwareID:   softwareUuid,
    Username:     username,
		LikedAt:      time.Now().UTC(),
	}, nil
}
