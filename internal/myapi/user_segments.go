package myapi

import (
	"encoding/json"
	"fmt"
	"github.com/gosimple/slug"
	"net/http"
)

type UserSegments struct {
	UserId      int      `json:"user_id"`
	AddSlugs    []string `json:"add_slugs"`
	DeleteSlugs []string `json:"delete_slugs"`
}

func (a *Api) UpdateUserSegments(w http.ResponseWriter, r *http.Request) {
	userSegments, err := GetUserSegmentsFromRequest(r)
	if err != nil {
		a.logger.With("path", r.URL.Path).Errorf("wrong segment: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user segments data"))
		return
	}
	err = a.logic.UpdateUserSegments(userSegments.UserId, userSegments.AddSlugs, userSegments.DeleteSlugs)
	if err != nil {
		a.logger.With("path", r.URL.Path).Errorf("updating user segments error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user segment data"))
		return
	}
	a.logger.With("path", r.URL.Path).
		Infof("successful user segments updating: userId: %d, addSlugs: %s, deleteSlugs: %s",
			userSegments.UserId, userSegments.AddSlugs, userSegments.DeleteSlugs)
	w.WriteHeader(http.StatusOK)
	return
}

func (a *Api) GetUserActiveSegments(w http.ResponseWriter, r *http.Request) {
	user, err := GetUserFromRequest(r)
	if err != nil {
		a.logger.With("path", r.URL.Path).Errorf("wrong segment: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user id data"))
		return
	}
	slugs, err := a.logic.GetUserActiveSegments(user.Id)
	if err != nil {
		a.logger.With("path", r.URL.Path).Errorf("getting user segments error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user id"))
		return
	}
	res, err := json.Marshal(slugs)
	if err != nil {
		a.logger.With("path", r.URL.Path).Errorf("json marshaling error: %w", err)
	}
	a.logger.With("path", r.URL.Path).
		Infof("successful user segments getting: userId %d, count: %d", user.Id, len(slugs))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	return
}

func GetUserSegmentsFromRequest(r *http.Request) (UserSegments, error) {
	var userSegments UserSegments
	err := json.NewDecoder(r.Body).Decode(&userSegments)
	if err != nil {
		return UserSegments{}, fmt.Errorf("json encoding: %w", err)
	}
	if userSegments.UserId == 0 {
		return UserSegments{}, fmt.Errorf("wrong user id")
	}
	for _, v := range userSegments.AddSlugs {
		if !slug.IsSlug(v) {
			return UserSegments{}, fmt.Errorf("wrong add slug field")
		}
	}
	for _, v := range userSegments.DeleteSlugs {
		if !slug.IsSlug(v) {
			return UserSegments{}, fmt.Errorf("wrong delete slug field")
		}
	}
	return userSegments, nil
}
