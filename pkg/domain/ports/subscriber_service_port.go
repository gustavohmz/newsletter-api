package ports

import domain "newsletter-app/pkg/domain/models"

type SubscriberServicePort interface {
	Subscribe(email string, category string) error
	Unsubscribe(email, category string) error
	GetSubscriberByEmail(email, category string) (*domain.Subscriber, error)
	GetSubscribers(email, category string, page, pageSize int) ([]domain.Subscriber, error)
}
