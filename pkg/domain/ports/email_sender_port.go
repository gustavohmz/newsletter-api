package ports

import domain "newsletter-app/pkg/domain/models"

type EmailSender interface {
	Send(subject, body string, to []string, attachments []*domain.Attachment) error
}
