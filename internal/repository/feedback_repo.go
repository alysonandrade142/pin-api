package repository

import (
	"pin-api/internal/entity"

	"gorm.io/gorm"
)

type FeedbackRepository interface {
	CreateFeedback(feedback *entity.Feedback) error
	GetFeedbackByID(id int) (*entity.Feedback, error)
	GetAllFeedbacks() ([]entity.Feedback, error)
	DeleteFeedback(id int) error
}

type feedbackRepository struct {
	db *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	return &feedbackRepository{db: db}
}

func (r *feedbackRepository) CreateFeedback(feedback *entity.Feedback) error {
	return r.db.Create(feedback).Error
}

func (r *feedbackRepository) GetFeedbackByID(id int) (*entity.Feedback, error) {
	var feedback entity.Feedback
	if err := r.db.First(&feedback, id).Error; err != nil {
		return nil, err
	}
	return &feedback, nil
}

func (r *feedbackRepository) GetAllFeedbacks() ([]entity.Feedback, error) {
	var feedbacks []entity.Feedback
	if err := r.db.Find(&feedbacks).Error; err != nil {
		return nil, err
	}
	return feedbacks, nil
}

func (r *feedbackRepository) DeleteFeedback(id int) error {
	return r.db.Delete(&entity.Feedback{}, id).Error
}
