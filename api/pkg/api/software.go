package api

import (
	"encoding/json"
	"fmt"
	"net/http"
  
  types "api/pkg/types"
)

func (s *APIServer) handleGetSoftware(w http.ResponseWriter, r *http.Request) error {
	software, err := s.store.GetSoftware()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, software)
}

func (s *APIServer) handleSoftware(w http.ResponseWriter, r *http.Request) error {
  if r.Method == "POST" {
    return s.handleCreateSoftware(w, r)
  } else if r.Method == "PUT" {
    return s.handleUpdateSoftware(w, r)
  } else if r.Method == "DELETE" {
    return s.handleDeleteSoftware(w, r)
  }

  return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetSoftwareByID(w http.ResponseWriter, r *http.Request) error {
  id := getID(r, "software-id")

  software, err := s.store.GetSoftwareByID(id)
  if err != nil {
    return err
  }

  return WriteJSON(w, http.StatusOK, software)
}

func (s *APIServer) handleCreateSoftware(w http.ResponseWriter, r *http.Request) error {
	req := new(types.CreateSoftwareRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
  
	software, err := types.NewSoftware(req.Name, req.Title, req.Description, req.Image, req.Url, req.UserID, req.Username)
	if err != nil {
		return err
	}
	if err := s.store.CreateSoftware(software); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, software)
}

func (s *APIServer) handleUpdateSoftware(w http.ResponseWriter, r *http.Request) error {
  return nil
}

func (s *APIServer) handleDeleteSoftware(w http.ResponseWriter, r *http.Request) error {
	id := getID(r, "software-id")

	if err := s.store.DeleteSoftware(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]string{"deleted": id})
}
