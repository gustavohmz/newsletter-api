package service_test

import (
	"testing"
	"time"

	domain "newsletter-app/pkg/domain/models"
	"newsletter-app/pkg/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSubscriberRepository struct {
	mock.Mock
}

func (m *MockSubscriberRepository) SaveSubscriber(subscriber domain.Subscriber) error {
	args := m.Called(subscriber)
	return args.Error(0)
}

func (m *MockSubscriberRepository) DeleteSubscriberByEmail(email, category string) error {
	args := m.Called(email, category)
	return args.Error(0)
}

func (m *MockSubscriberRepository) GetSubscriberByEmailAndCategory(email, category string) (*domain.Subscriber, error) {
	args := m.Called(email, category)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.Subscriber), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockSubscriberRepository) GetSubscribers(email, category string, page, pageSize int) ([]domain.Subscriber, error) {
	args := m.Called(email, category, page, pageSize)
	if args.Get(0) != nil {
		return args.Get(0).([]domain.Subscriber), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestSubscribe(t *testing.T) {
	mockRepo := new(MockSubscriberRepository)
	subscriberService := service.NewSubscriberService(mockRepo)

	subscriber := domain.Subscriber{
		Email:            "test@example.com",
		SubscriptionDate: time.Now(),
		Category:         "Tech",
	}

	mockRepo.On("SaveSubscriber", subscriber).Return(nil)

	err := subscriberService.Subscribe(subscriber.Email, subscriber.Category)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUnsubscribe(t *testing.T) {
	mockRepo := new(MockSubscriberRepository)
	subscriberService := service.NewSubscriberService(mockRepo)

	mockRepo.On("DeleteSubscriberByEmail", "test@example.com", "Tech").Return(nil)

	err := subscriberService.Unsubscribe("test@example.com", "Tech")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetSubscriberByEmail(t *testing.T) {
	mockRepo := new(MockSubscriberRepository)
	subscriberService := service.NewSubscriberService(mockRepo)

	subscriber := &domain.Subscriber{
		Email:            "test@example.com",
		SubscriptionDate: time.Now(),
		Category:         "Tech",
	}

	mockRepo.On("GetSubscriberByEmailAndCategory", "test@example.com", "Tech").Return(subscriber, nil)

	result, err := subscriberService.GetSubscriberByEmail("test@example.com", "Tech")
	assert.NoError(t, err)
	assert.Equal(t, subscriber, result)
	mockRepo.AssertExpectations(t)
}

func TestGetSubscribers(t *testing.T) {
	mockRepo := new(MockSubscriberRepository)
	subscriberService := service.NewSubscriberService(mockRepo)

	subscribers := []domain.Subscriber{
		{Email: "test1@example.com", Category: "Tech"},
		{Email: "test2@example.com", Category: "Tech"},
	}

	mockRepo.On("GetSubscribers", "", "Tech", 1, 10).Return(subscribers, nil)

	result, err := subscriberService.GetSubscribers("", "Tech", 1, 10)
	assert.NoError(t, err)
	assert.Equal(t, subscribers, result)
	mockRepo.AssertExpectations(t)
}
