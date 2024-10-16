package service

import (
	"pin-api/internal/dto"
	"pin-api/internal/entity"
	"pin-api/internal/repository"
)

type FeedbackService interface {
	CreateFeedback(feedbackDTO dto.FeedbackDTO) error
	GetFeedbackByID(id int) (*entity.Feedback, error)
	GetAllFeedbacks() ([]entity.Feedback, error)
	DeleteFeedback(id int) error
}

type feedbackService struct {
	repo repository.FeedbackRepository
}

func NewFeedbackService(repo repository.FeedbackRepository) FeedbackService {
	return &feedbackService{repo: repo}
}

func (s *feedbackService) CreateFeedback(feedbackDTO dto.FeedbackDTO) error {
	feedback := entity.Feedback{
		FuncionarioID:                     feedbackDTO.FuncionarioID,
		Feedback:                          feedbackDTO.Feedback,
		ComentariosFeedback:               feedbackDTO.ComentariosFeedback,
		InteracaoGestor:                   feedbackDTO.InteracaoGestor,
		ComentariosInteracaoGestor:        feedbackDTO.ComentariosInteracaoGestor,
		ClarezaCarreira:                   feedbackDTO.ClarezaCarreira,
		ComentariosClarezaCarreira:        feedbackDTO.ComentariosClarezaCarreira,
		ExpectativaPermanencia:            feedbackDTO.ExpectativaPermanencia,
		ComentariosExpectativaPermanencia: feedbackDTO.ComentariosExpectativaPermanencia,
		ENPS:                              feedbackDTO.ENPS,
		AbertaENPS:                        feedbackDTO.AbertaENPS,
	}
	return s.repo.CreateFeedback(&feedback)
}

func (s *feedbackService) GetFeedbackByID(id int) (*entity.Feedback, error) {
	return s.repo.GetFeedbackByID(id)
}

func (s *feedbackService) GetAllFeedbacks() ([]entity.Feedback, error) {
	return s.repo.GetAllFeedbacks()
}

func (s *feedbackService) DeleteFeedback(id int) error {
	return s.repo.DeleteFeedback(id)
}
