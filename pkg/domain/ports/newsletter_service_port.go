package ports

import (
	"net/http"
	domain "newsletter-app/pkg/domain/models"
	"newsletter-app/pkg/infrastructure/adapters/email"
	"newsletter-app/pkg/service/Dtos/request"
)

type NewsletterServicePort interface {
	SaveNewsletter(newsletter domain.Newsletter) error
	GetNewsletterByCategory(category string) (*domain.Newsletter, error)
	GetNewsletterByID(newsletterID string) (*domain.Newsletter, error)
	GetNewsletters(searchName string, page int, pageSize int) ([]domain.Newsletter, error)
	SendNewsletter(w http.ResponseWriter, r *http.Request, newsletterID string, emailSender email.EmailSender) error
	UpdateNewsletter(updateRequest request.UpdateNewsletterRequest) error
	DeleteNewsletter(id string) error
}
