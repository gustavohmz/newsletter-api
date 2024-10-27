package ports

import domain "newsletter-app/pkg/domain/models"

type NewsletterRepositoryPort interface {
	SaveNewsletter(newsletter domain.Newsletter) error
	GetNewsletterByCategory(category string) (*domain.Newsletter, error)
	GetNewsletterByID(newsletterID string) (*domain.Newsletter, error)
	GetNewsletters(searchName string, page int, pageSize int) ([]domain.Newsletter, error)
	UpdateNewsletter(newsletter domain.Newsletter) error
	DeleteNewsletterByID(id string) error
}
