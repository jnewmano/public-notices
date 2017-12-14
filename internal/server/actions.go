package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UpcomingNotice struct {
	Date      time.Time
	Entity    string
	Longitude float64
	Latitude  float64
}

func (s *Server) UpcomingActionHandler(w http.ResponseWriter, r *http.Request) {
	// returns a json object array with a list of upcoming
	// planning commission agenda items

	entity := r.URL.Query().Get("entity")

	fmt.Println("loading notices for", entity)

	m := s.processor.Meeting(entity)

	_ = json.NewEncoder(w).Encode(m)

}
