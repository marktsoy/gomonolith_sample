package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/marktsoy/gomonolith_sample/internal/app/models"
	"github.com/marktsoy/gomonolith_sample/internal/app/utils"
)

func (s *Server) createBundle() func(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Size     int    `json:"size"`
		Priority string `json:"priority"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		req := &request{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		bundle := &models.Bundle{
			Size: req.Size,
		}
		switch req.Priority {
		default:
			bundle.Priority = models.PriorityLow
		case "medium":
			bundle.Priority = models.PriorityMedium
		case "high":
			bundle.Priority = models.PriorityHigh
		}
		repo := s.store.Bundle()
		repo.Create(bundle)

		if err := json.NewEncoder(w).Encode(bundle); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(200)
		go func() {
			for i := 0; i < bundle.Size; i++ {
				m := &models.Message{
					Content:  "Random Content " + strconv.Itoa(i) + " " + strconv.Itoa(bundle.ID) + " Priority " + strconv.Itoa(bundle.Priority),
					Priority: bundle.Priority,
					BundleID: bundle.ID,
				}
				s.store.Message().Create(m)
				s.msgs <- &utils.Action{
					Name:    "messageCreated",
					Payload: m,
				}
			}
		}()
		return
	}
}
