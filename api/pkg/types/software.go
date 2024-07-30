package types

import (
	"time"

	"github.com/google/uuid"
)

type SoftwareRequest struct {
	ID      uuid.UUID   `json:"id"`
}

type CreateSoftwareRequest struct {
  Name          string    `json:"name"`
  Title         string    `json:"title"`
  Description   string    `json:"description"`
  Image         string    `json:"image"`
  Url           string    `json:"url"`
  Username      string    `json:"username"`
}

type Software struct {
	ID            uuid.UUID `json:"id"`
  Name          string    `json:"name"`
  Title         string    `json:"title"`
  Description   string    `json:"description"`
  Image         string    `json:"image"`
  Url           string    `json:"url"`
  Username      string    `json:"username"`
	CreatedAt     time.Time `json:"createdAt"`
}

func NewSoftware(name, title, description, image, url, username string) (*Software, error) {
	return &Software{
    Name:              name,
    Title:             title,
    Description:       description,
    Image:             image,
    Url:               url,
    Username:          username,
		CreatedAt:         time.Now().UTC(),
	}, nil
}
