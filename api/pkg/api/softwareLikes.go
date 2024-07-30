package api

import (
	"encoding/json"
	"fmt"
	"net/http"
  
  types "api/pkg/types"
)

func (s *APIServer) handleGetSoftwareLikesByID(w http.ResponseWriter, r *http.Request) error {
  id := getID(r, "software-id")

  likes, err := s.store.GetSoftwareLikesBySoftware(id)
  if err != nil {
    return err
  }

  return WriteJSON(w, http.StatusOK, likes)
}

func (s *APIServer) handleSoftwareLike(w http.ResponseWriter, r *http.Request) error {
  if r.Method == "POST" {
    return s.handleCreateSoftwareLike(w, r)
  } else if r.Method == "DELETE" {
    return s.handleDeleteSoftwareLike(w, r)
  }

  return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleCreateSoftwareLike(w http.ResponseWriter, r *http.Request) error {
	req := new(types.CreateSoftwareLikeRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

  like, err := types.NewSoftwareLike(req.SoftwareID, req.UserID, req.Username)
	if err != nil {
		return err
	}

	if err = s.store.CreateSoftwareLike(like); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, like)
}

func (s *APIServer) handleDeleteSoftwareLike(w http.ResponseWriter, r *http.Request) error {
  softwareId := getID(r, "software-id")
	username := getID(r, "username")

	if err := s.store.DeleteSoftwareLike(softwareId, username); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]string{"like deleted": username})
}
