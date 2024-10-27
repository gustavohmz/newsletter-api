package service_test

import (
	domain "newsletter-app/pkg/domain/models"
	"newsletter-app/pkg/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockNewsletterRepository struct {
	mock.Mock
}

func (m *MockNewsletterRepository) SaveNewsletter(newsletter domain.Newsletter) error {
	args := m.Called(newsletter)
	return args.Error(0)
}

func (m *MockNewsletterRepository) GetNewsletterByCategory(category string) (*domain.Newsletter, error) {
	args := m.Called(category)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Newsletter), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockNewsletterRepository) GetNewsletterByID(newsletterID string) (*domain.Newsletter, error) {
	args := m.Called(newsletterID)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Newsletter), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockNewsletterRepository) GetNewsletters(searchName string, page int, pageSize int) ([]domain.Newsletter, error) {
	args := m.Called(searchName, page, pageSize)
	return args.Get(0).([]domain.Newsletter), args.Error(1)
}

func (m *MockNewsletterRepository) UpdateNewsletter(newsletter domain.Newsletter) error {
	args := m.Called(newsletter)
	return args.Error(0)
}

func (m *MockNewsletterRepository) DeleteNewsletterByID(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockSubscriberRepository) GetSubscribersByCategory(category string) ([]domain.Subscriber, error) {
	args := m.Called(category)
	return args.Get(0).([]domain.Subscriber), args.Error(1)
}

type MockEmailSender struct {
	mock.Mock
}

func (m *MockEmailSender) Send(subject, body string, to []string, attachments []*domain.Attachment) error {
	args := m.Called(subject, body, to, attachments)
	return args.Error(0)
}

func TestSaveNewsletter(t *testing.T) {
	mockNewsletterRepo := new(MockNewsletterRepository)
	newsletterService := service.NewNewsletterService(mockNewsletterRepo, new(MockSubscriberRepository))

	newNewsletter := domain.Newsletter{
		ID:       primitive.NewObjectID(),
		Name:     "Test Newsletter",
		Category: "Tech",
	}

	mockNewsletterRepo.On("SaveNewsletter", newNewsletter).Return(nil)

	err := newsletterService.SaveNewsletter(newNewsletter)
	assert.NoError(t, err)
	mockNewsletterRepo.AssertExpectations(t)
}

func TestGetNewsletterByCategory(t *testing.T) {
	mockNewsletterRepo := new(MockNewsletterRepository)
	newsletterService := service.NewNewsletterService(mockNewsletterRepo, new(MockSubscriberRepository))

	mockNewsletter := &domain.Newsletter{ID: primitive.NewObjectID(), Category: "Tech"}
	mockNewsletterRepo.On("GetNewsletterByCategory", "Tech").Return(mockNewsletter, nil)

	result, err := newsletterService.GetNewsletterByCategory("Tech")
	assert.NoError(t, err)
	assert.Equal(t, mockNewsletter, result)
	mockNewsletterRepo.AssertExpectations(t)
}

func TestGetNewsletterByID(t *testing.T) {
	mockNewsletterRepo := new(MockNewsletterRepository)
	newsletterService := service.NewNewsletterService(mockNewsletterRepo, new(MockSubscriberRepository))

	mockNewsletter := &domain.Newsletter{ID: primitive.NewObjectID(), Category: "Tech"}
	mockNewsletterRepo.On("GetNewsletterByID", mockNewsletter.ID.Hex()).Return(mockNewsletter, nil)

	result, err := newsletterService.GetNewsletterByID(mockNewsletter.ID.Hex())
	assert.NoError(t, err)
	assert.Equal(t, mockNewsletter, result)
	mockNewsletterRepo.AssertExpectations(t)
}

func TestGetNewsletters(t *testing.T) {
	mockNewsletterRepo := new(MockNewsletterRepository)
	newsletterService := service.NewNewsletterService(mockNewsletterRepo, new(MockSubscriberRepository))

	newsletters := []domain.Newsletter{
		{ID: primitive.NewObjectID(), Category: "Tech"},
		{ID: primitive.NewObjectID(), Category: "Science"},
	}

	mockNewsletterRepo.On("GetNewsletters", "", 1, 10).Return(newsletters, nil)

	result, err := newsletterService.GetNewsletters("", 1, 10)
	assert.NoError(t, err)
	assert.Equal(t, newsletters, result)
	mockNewsletterRepo.AssertExpectations(t)
}

func TestDeleteNewsletter(t *testing.T) {
	mockNewsletterRepo := new(MockNewsletterRepository)
	newsletterService := service.NewNewsletterService(mockNewsletterRepo, new(MockSubscriberRepository))

	mockNewsletterRepo.On("DeleteNewsletterByID", "1").Return(nil)

	err := newsletterService.DeleteNewsletter("1")
	assert.NoError(t, err)
	mockNewsletterRepo.AssertExpectations(t)
}
