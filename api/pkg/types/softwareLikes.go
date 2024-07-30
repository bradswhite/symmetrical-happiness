package types

import (
	"time"

	"github.com/google/uuid"
)

type SoftwareLikeRequest struct {
	SoftwareID   uuid.UUID   `json:"softwareId"`
	UserID       uuid.UUID   `json:"userId"`
	Username     string      `json:"username"`
}

type CreateSoftwareLikeRequest struct {
	SoftwareID   uuid.UUID   `json:"softwareId"`
	UserID       uuid.UUID   `json:"userId"`
	Username     string      `json:"username"`
}

type SoftwareLike struct {
	SoftwareID   uuid.UUID   `json:"softwareId"`
	UserID       uuid.UUID   `json:"userId"`
	Username     string      `json:"username"`
	LikedAt      time.Time   `json:"likedAt"`
}

func NewSoftwareLike(softwareId, userId uuid.UUID, username string) (*SoftwareLike, error) {
	return &SoftwareLike{
    SoftwareID:   softwareId,
    UserID:       userId,
    Username:     username,
		LikedAt:      time.Now().UTC(),
	}, nil
}
