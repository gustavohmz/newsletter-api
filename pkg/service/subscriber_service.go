package service

import (
	domain "newsletter-app/pkg/domain/models"
	"newsletter-app/pkg/domain/ports"
	"time"
)

var _ ports.SubscriberServicePort = (*SubscriberServiceImpl)(nil)

type SubscriberServiceImpl struct {
	subscriberRepository ports.SubscriberRepositoryPort
}

func NewSubscriberService(subscriberRepo ports.SubscriberRepositoryPort) ports.SubscriberServicePort {
	return &SubscriberServiceImpl{
		subscriberRepository: subscriberRepo,
	}
}

func (s *SubscriberServiceImpl) Subscribe(email string, category string) error {
	subscriber := domain.Subscriber{
		Email:            email,
		SubscriptionDate: time.Now(),
		Category:         category,
	}
	return s.subscriberRepository.SaveSubscriber(subscriber)
}

func (s *SubscriberServiceImpl) Unsubscribe(email, category string) error {
	return s.subscriberRepository.DeleteSubscriberByEmail(email, category)
}

func (s *SubscriberServiceImpl) GetSubscriberByEmail(email, category string) (*domain.Subscriber, error) {
	return s.subscriberRepository.GetSubscriberByEmailAndCategory(email, category)
}

func (s *SubscriberServiceImpl) GetSubscribers(email, category string, page, pageSize int) ([]domain.Subscriber, error) {
	return s.subscriberRepository.GetSubscribers(email, category, page, pageSize)
}
