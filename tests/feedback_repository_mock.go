package tests

import (
	"pin-api/internal/entity"

	"github.com/stretchr/testify/mock"
)

type FeedbackRepositoryMock struct {
	mock.Mock
}

func (m *FeedbackRepositoryMock) GetFeedbackByID(id int) (*entity.Feedback, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Feedback), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *FeedbackRepositoryMock) GetAllFeedbacks() ([]entity.Feedback, error) {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]entity.Feedback), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *FeedbackRepositoryMock) CreateFeedback(feedback *entity.Feedback) error {
	args := m.Called(feedback)
	return args.Error(0)
}

func (m *FeedbackRepositoryMock) DeleteFeedback(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
