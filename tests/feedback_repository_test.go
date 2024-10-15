package tests

import (
	"pin-api/internal/entity"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFeedbackByID(t *testing.T) {
	// Crie uma instância do mock
	feedbackRepoMock := new(FeedbackRepositoryMock)

	// Crie um feedback para teste
	feedback := &entity.Feedback{
		ID:            1,
		FuncionarioID: 1,
		Feedback:      "Bom trabalho!",
	}

	// Configure o mock para retornar o feedback esperado
	feedbackRepoMock.On("GetFeedbackByID", feedback.ID).Return(feedback, nil)

	// Teste a busca usando o mock
	result, err := feedbackRepoMock.GetFeedbackByID(feedback.ID)
	assert.NoError(t, err)
	assert.Equal(t, feedback.Feedback, result.Feedback)

	// Verifique se o método foi chamado
	feedbackRepoMock.AssertExpectations(t)
}

func TestFindAllFeedbacks(t *testing.T) {
	// Crie uma instância do mock
	feedbackRepoMock := new(FeedbackRepositoryMock)

	// Crie um feedback para teste
	feedback := &entity.Feedback{
		ID:            1,
		FuncionarioID: 1,
		Feedback:      "Bom trabalho!",
	}

	// Configure o mock para retornar o feedback esperado
	feedbackRepoMock.On("GetAllFeedbacks").Return([]entity.Feedback{*feedback}, nil)

	// Teste a busca usando o mock
	result, err := feedbackRepoMock.GetAllFeedbacks()
	assert.NoError(t, err)
	assert.Equal(t, feedback.Feedback, result[0].Feedback)

	// Verifique se o método foi chamado
	feedbackRepoMock.AssertExpectations(t)
}

func TestCreateFeedback(t *testing.T) {
	// Crie uma instância do mock
	feedbackRepoMock := new(FeedbackRepositoryMock)

	// Crie um feedback para teste
	feedback := &entity.Feedback{
		ID:            1,
		FuncionarioID: 1,
		Feedback:      "Bom trabalho!",
	}

	// Configure o mock para retornar o feedback esperado
	feedbackRepoMock.On("CreateFeedback", feedback).Return(nil)

	// Teste a busca usando o mock
	err := feedbackRepoMock.CreateFeedback(feedback)
	assert.NoError(t, err)

	// Verifique se o método foi chamado
	feedbackRepoMock.AssertExpectations(t)
}
