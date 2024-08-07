package subscription

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

type SubscribeReqBody struct {
	UserID       string `json:"userID"`
	Plan         string `json:"plan"`
	DurationDays int    `json:"durationDays"`
}

type Handler struct {
	s *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) Subscribe(w http.ResponseWriter, r *http.Request) {
	var reqBody SubscribeReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Debug().Err(err).Msg("Failed to decode JSON")
		return
	}

	h.s.SubscribeUser(reqBody.UserID, reqBody.Plan, reqBody.DurationDays)
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) ListSubscriptions(w http.ResponseWriter, r *http.Request) {
	subscriptions := h.s.GetSubscriptions()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subscriptions)
}

type GetUserSubscriptionsReqParams struct {
	UserID string
}

func (h *Handler) GetUserSubscriptions(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		http.Error(w, "Missing user_id parameter", http.StatusBadRequest)
		return
	}

	subscriptions := h.s.GetSubscriptionByUserID(userID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subscriptions)
}
