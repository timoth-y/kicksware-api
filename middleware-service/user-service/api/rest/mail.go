package rest

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *handler) SendEmailConfirmation(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	userID := query.Get("userID")
	callbackURL := query.Get("callbackURL")
	err := h.mail.SendEmailConfirmation(userID, callbackURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *handler) SendResetPassword(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r,"userID")
	callbackURL := chi.URLParam(r,"callbackURL")
	err := h.mail.SendResetPassword(userID, callbackURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}

func (h *handler) SendNotification(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r,"userID")
	callbackURL := chi.URLParam(r,"callbackURL")
	err := h.mail.SendNotification(userID, callbackURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.setupResponse(w, nil, http.StatusOK)
}