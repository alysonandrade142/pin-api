package handler

import (
	"encoding/json"
	"net/http"
	"pin-api/internal/dto"
	"pin-api/internal/service"
	"strconv"

	"github.com/gorilla/mux"
)

type FeedbackHandler struct {
	service service.FeedbackService
}

func NewFeedbackHandler(service service.FeedbackService) *FeedbackHandler {
	return &FeedbackHandler{service: service}
}

func (h *FeedbackHandler) CreateFeedback(w http.ResponseWriter, r *http.Request) {
	var feedbackDTO dto.FeedbackDTO
	if err := json.NewDecoder(r.Body).Decode(&feedbackDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateFeedback(feedbackDTO); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Feedback criado com sucesso!")
}

func (h *FeedbackHandler) GetFeedbackByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	feedback, err := h.service.GetFeedbackByID(id)
	if err != nil {
		http.Error(w, "Feedback not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feedback)
}

func (h *FeedbackHandler) GetAllFeedbacks(w http.ResponseWriter, r *http.Request) {
	feedbacks, err := h.service.GetAllFeedbacks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(feedbacks)
}

func (h *FeedbackHandler) DeleteFeedback(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteFeedback(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Feedback deletado com sucesso!")
}
