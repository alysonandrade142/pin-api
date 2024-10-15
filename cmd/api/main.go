package main

import (
	"net/http"
	"pin-api/internal/handler"
	"pin-api/internal/repository"
	"pin-api/internal/service"
	"pin-api/pkg/database"
	"pin-api/pkg/logger"

	"github.com/gorilla/mux"
)

func main() {
	// Configura o logger
	log := logger.SetupLogger()

	// Inicializa o banco de dados
	db := database.InitDB()

	// Inicializa o repositório e o serviço para Feedback
	feedbackRepo := repository.NewFeedbackRepository(db)
	feedbackService := service.NewFeedbackService(feedbackRepo)
	feedbackHandler := handler.NewFeedbackHandler(feedbackService)

	// Roteador mux
	r := mux.NewRouter()

	// Rotas da API para feedback
	r.HandleFunc("/feedback", feedbackHandler.CreateFeedback).Methods("POST")
	r.HandleFunc("/feedback/{id}", feedbackHandler.GetFeedbackByID).Methods("GET")
	r.HandleFunc("/feedback", feedbackHandler.GetAllFeedbacks).Methods("GET")

	// Iniciar o servidor
	log.Println("Servidor rodando na porta 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Falha ao iniciar o servidor: %v", err)
	}
}
