package ports

import domain "newsletter-app/pkg/domain/models"

type SubscriberRepositoryPort interface {
	SaveSubscriber(subscriber domain.Subscriber) error
	DeleteSubscriberByEmail(email, category string) error
	GetSubscriberByEmailAndCategory(email, category string) (*domain.Subscriber, error)
	GetSubscribers(email, category string, page, pageSize int) ([]domain.Subscriber, error)
	GetSubscribersByCategory(category string) ([]domain.Subscriber, error)
}
