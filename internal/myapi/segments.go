package myapi

import (
	"encoding/json"
	"fmt"
	"github.com/gosimple/slug"
	"net/http"
)

type Segment struct {
	Slug    string `json:"slug"`
	Percent int    `json:"percent"`
}

func (a *Api) AddSegment(w http.ResponseWriter, r *http.Request) {
	segment, err := GetSegmentFromRequest(r)
	if err != nil {
		a.logger.With("path", r.URL.Path).Errorf("wrong segment: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid segment"))
		return
	}
	err = a.logic.AddSegment(segment.Slug)
	if err != nil {
		a.logger.With("path", r.URL.Path).Errorf("add segment error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid segment"))
		return
	}
	if segment.Percent > 0 {
		if err = a.logic.AddSegmentForUsersPercent(segment.Slug, segment.Percent); err != nil {
			a.logger.With("path", r.URL.Path).Errorf("add segment for users error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Segment created successfully, but no user implements it"))
			return
		}
	}
	a.logger.With("path", r.URL.Path).Infof("successful segment addition: slug: %s", segment.Slug)
	w.WriteHeader(http.StatusOK)
	return
}

func (a *Api) RemoveSegment(w http.ResponseWriter, r *http.Request) {
	segment, err := GetSegmentFromRequest(r)
	if err != nil {
		a.logger.With("path", r.URL.Path).Errorf("wrong segment: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid segment"))
		return
	}
	err = a.logic.RemoveSegment(segment.Slug)
	if err != nil {
		a.logger.With("path", r.URL.Path).Errorf("remove segment error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid segment"))
		return
	}
	a.logger.With("path", r.URL.Path).Infof("successful segment removal: slug: %s", segment.Slug)
	w.WriteHeader(http.StatusOK)
	return
}

func GetSegmentFromRequest(r *http.Request) (Segment, error) {
	var segment Segment
	err := json.NewDecoder(r.Body).Decode(&segment)
	if err != nil {
		return Segment{}, fmt.Errorf("json encoding: %w", err)
	}
	if !slug.IsSlug(segment.Slug) {
		return Segment{}, fmt.Errorf("wrong slug field")
	}
	return segment, nil
}
