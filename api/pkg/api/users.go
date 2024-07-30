package api

import (
	"encoding/json"
	"fmt"
	"net/http"
  
  types "api/pkg/types"
)

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	var req types.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	acc, err := s.store.GetUserByEmail(req.Email)
	if err != nil {
		return err
	}

	if !acc.ValidPassword(req.Password) {
		return fmt.Errorf("not authenticated")
	}

	token, err := createJWT(acc)
	if err != nil {
		return err
	}

	resp := types.LoginResponse{
    ID: acc.ID,
		Token:  token,
		Username: acc.Username,
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleSignUp(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
    req := new(types.CreateUserRequest)
    if err := json.NewDecoder(r.Body).Decode(req); err != nil {
      return err
    }

    user, err := types.NewUser(req.Email, req.Username, req.Password)
    if err != nil {
      return err
    }
    if err := s.store.CreateUser(user); err != nil {
      return err
    }

    return WriteJSON(w, http.StatusOK, user)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id := getID(r, "username")

	if err := s.store.DeleteUser(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]string{"deleted": id})
}
