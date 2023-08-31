package myapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id int `json:"id"`
}

func (a *Api) AddUser(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserFromRequest(r)
	if err != nil {
		a.logger.With("path", r.URL.Path).Errorf("wrong user: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user"))
		return
	}
	err = a.logic.AddUser(user.Id)
	if err != nil {
		a.logger.With("path", r.URL.Path).Errorf("add user error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user"))
		return
	}
	a.logger.With("path", r.URL.Path).Infof("successful user addition: slug: %d", user.Id)
	w.WriteHeader(http.StatusOK)
	return
}

func (a *Api) RemoveUser(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserFromRequest(r)
	if err != nil {
		a.logger.With("path", r.URL.Path).Errorf("wrong user: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user"))
		return
	}
	err = a.logic.RemoveUser(user.Id)
	if err != nil {
		a.logger.With("path", r.URL.Path).Errorf("remove user error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user"))
		return
	}
	a.logger.With("path", r.URL.Path).Infof("successful user removal: slug: %d", user.Id)
	w.WriteHeader(http.StatusOK)
	return
}

func GetUserFromRequest(r *http.Request) (User, error) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return User{}, fmt.Errorf("json encoding: %w", err)
	}
	if user.Id == 0 {
		return User{}, fmt.Errorf("wrong id field")
	}
	return user, nil
}
