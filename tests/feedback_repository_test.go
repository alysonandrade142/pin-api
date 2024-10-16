package tests

import (
	"errors"
	"pin-api/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFeedbackByID_Success(t *testing.T) {
	feedbackRepoMock := new(FeedbackRepositoryMock)

	feedback := &entity.Feedback{
		ID:            1,
		FuncionarioID: 1,
		Feedback:      "Bom trabalho!",
	}

	feedbackRepoMock.On("GetFeedbackByID", feedback.ID).Return(feedback, nil)

	result, err := feedbackRepoMock.GetFeedbackByID(feedback.ID)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, feedback.Feedback, result.Feedback)

	feedbackRepoMock.AssertExpectations(t)
}

func TestFindFeedbackByID_Error(t *testing.T) {
	feedbackRepoMock := new(FeedbackRepositoryMock)

	feedbackRepoMock.On("GetFeedbackByID", 2).Return(nil, errors.New("Feedback não encontrado"))

	result, err := feedbackRepoMock.GetFeedbackByID(2)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "Feedback não encontrado")

	feedbackRepoMock.AssertExpectations(t)
}

func TestFindAllFeedbacks_Success(t *testing.T) {
	feedbackRepoMock := new(FeedbackRepositoryMock)

	feedback := &entity.Feedback{
		ID:            1,
		FuncionarioID: 1,
		Feedback:      "Bom trabalho!",
	}

	feedbackRepoMock.On("GetAllFeedbacks").Return([]entity.Feedback{*feedback}, nil)

	result, err := feedbackRepoMock.GetAllFeedbacks()
	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, feedback.Feedback, result[0].Feedback)

	feedbackRepoMock.AssertExpectations(t)
}

func TestFindAllFeedbacks_Error(t *testing.T) {
	feedbackRepoMock := new(FeedbackRepositoryMock)

	feedbackRepoMock.On("GetAllFeedbacks").Return(nil, errors.New("Erro ao buscar feedbacks"))

	result, err := feedbackRepoMock.GetAllFeedbacks()
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "Erro ao buscar feedbacks")

	feedbackRepoMock.AssertExpectations(t)
}

func TestCreateFeedback_Success(t *testing.T) {
	feedbackRepoMock := new(FeedbackRepositoryMock)

	feedback := &entity.Feedback{
		ID:            1,
		FuncionarioID: 1,
		Feedback:      "Bom trabalho!",
	}

	feedbackRepoMock.On("CreateFeedback", feedback).Return(nil)

	err := feedbackRepoMock.CreateFeedback(feedback)
	assert.NoError(t, err)

	feedbackRepoMock.AssertExpectations(t)
}

func TestCreateFeedback_Error(t *testing.T) {
	feedbackRepoMock := new(FeedbackRepositoryMock)

	feedback := &entity.Feedback{
		ID:            2,
		FuncionarioID: 2,
		Feedback:      "Necessita melhorar.",
	}

	feedbackRepoMock.On("CreateFeedback", feedback).Return(errors.New("Erro ao criar feedback"))

	err := feedbackRepoMock.CreateFeedback(feedback)
	assert.Error(t, err)
	assert.EqualError(t, err, "Erro ao criar feedback")

	feedbackRepoMock.AssertExpectations(t)
}

func TestDeleteFeedback_Success(t *testing.T) {
	feedbackRepoMock := new(FeedbackRepositoryMock)

	feedback := &entity.Feedback{
		ID: 1,
	}

	feedbackRepoMock.On("DeleteFeedback", feedback.ID).Return(nil)

	err := feedbackRepoMock.DeleteFeedback(feedback.ID)
	assert.NoError(t, err)

	feedbackRepoMock.AssertExpectations(t)
}
