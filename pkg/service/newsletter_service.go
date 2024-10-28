package service

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	domain "newsletter-app/pkg/domain/models"
	"newsletter-app/pkg/domain/ports"
	"newsletter-app/pkg/infrastructure/adapters/email"
	"newsletter-app/pkg/service/Dtos/request"
	"strings"
)

var _ ports.NewsletterServicePort = (*NewsletterService)(nil)

type NewsletterService struct {
	newsletterRepository ports.NewsletterRepositoryPort
	subscriberRepository ports.SubscriberRepositoryPort
}

func NewNewsletterService(
	newsletterRepo ports.NewsletterRepositoryPort,
	subscriberRepo ports.SubscriberRepositoryPort,
) *NewsletterService {
	return &NewsletterService{
		newsletterRepository: newsletterRepo,
		subscriberRepository: subscriberRepo,
	}
}

func (s *NewsletterService) SaveNewsletter(newsletter domain.Newsletter) error {
	var decodedAttachments []domain.Attachment

	for _, base64Attachment := range newsletter.Attachments {
		attachment := domain.Attachment{
			Name: base64Attachment.Name,
			Data: base64Attachment.Data,
			Type: base64Attachment.Type,
		}
		decodedAttachments = append(decodedAttachments, attachment)
	}

	newsletter.Attachments = decodedAttachments
	return s.newsletterRepository.SaveNewsletter(newsletter)
}

func (s *NewsletterService) GetNewsletterByCategory(category string) (*domain.Newsletter, error) {
	return s.newsletterRepository.GetNewsletterByCategory(category)
}

func (s *NewsletterService) GetNewsletterByID(newsletterID string) (*domain.Newsletter, error) {
	return s.newsletterRepository.GetNewsletterByID(newsletterID)
}

func (s *NewsletterService) GetNewsletters(searchName string, page int, pageSize int) ([]domain.Newsletter, error) {
	newsletters, err := s.newsletterRepository.GetNewsletters(searchName, page, pageSize)
	if err != nil {
		return nil, err
	}

	return newsletters, nil
}

func (s *NewsletterService) SendNewsletter(w http.ResponseWriter, r *http.Request, newsletterID string, emailSender email.EmailSender) error {
	newsletter, err := s.GetNewsletterByID(newsletterID)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve newsletter")
		return err
	}

	subscribers, err := s.subscriberRepository.GetSubscribersByCategory(newsletter.Category)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve subscribers")
		return err
	}

	if len(subscribers) == 0 {
		RespondWithError(w, http.StatusBadRequest, "No subscribers to send the newsletter to")
		return nil
	}

	decodedAttachments, err := DecodeAttachments(newsletter.Attachments)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Failed to decode attachments")
		return err
	}

	for _, subscriber := range subscribers {
		fmt.Printf("Subscriber: %+v\n", subscriber)

		if newsletter.Content == "" {
			RespondWithError(w, http.StatusBadRequest, "Newsletter content is empty")
			return nil
		}

		emailCategoryConcatenation := fmt.Sprintf("%s|%s", subscriber.Email, subscriber.Category)
		newsletterContent := strings.ReplaceAll(newsletter.Content, "{email}", emailCategoryConcatenation)
		urlUnsubscribe := strings.ReplaceAll(newsletterContent, "{hostDomain}", "http://localhost:4200/")

		err = emailSender.Send(newsletter.Subject, urlUnsubscribe, []string{subscriber.Email}, decodedAttachments)
		if err != nil {
			fmt.Printf("Error sending newsletter to %s: %s\n", subscriber.Email, err.Error())
			continue
		}

		fmt.Printf("Newsletter sent to %s\n", subscriber.Email)
	}

	RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"status":  "OK",
		"message": "Newsletter sent successfully",
	})
	return nil
}

func DecodeAttachments(attachments []domain.Attachment) ([]*domain.Attachment, error) {
	var decodedAttachments []*domain.Attachment

	for _, attachment := range attachments {
		_, err := base64.StdEncoding.DecodeString(attachment.Data)
		if err != nil {
			return nil, errors.New("failed to decode attachment data")
		}

		decodedAttachment := &domain.Attachment{
			Name: attachment.Name,
			Data: attachment.Data,
			Type: attachment.Type,
		}

		decodedAttachments = append(decodedAttachments, decodedAttachment)
	}

	return decodedAttachments, nil
}

func (s *NewsletterService) UpdateNewsletter(updateRequest request.UpdateNewsletterRequest) error {
	if updateRequest.ID.IsZero() {
		return errors.New("ID is required for update")
	}

	existingNewsletter, err := s.GetNewsletterByID(updateRequest.ID.Hex())
	if err != nil {
		return err
	}

	existingNewsletter.Name = updateRequest.Name
	existingNewsletter.Category = updateRequest.Category
	existingNewsletter.Subject = updateRequest.Subject
	existingNewsletter.Content = updateRequest.Content

	if len(updateRequest.Attachments) > 0 {
		existingNewsletter.Attachments = make([]domain.Attachment, len(updateRequest.Attachments))
		for i, attachment := range updateRequest.Attachments {
			existingNewsletter.Attachments[i] = domain.Attachment{
				Name: attachment.Name,
				Data: attachment.Data,
				Type: attachment.Type,
			}
		}
	} else {
		existingNewsletter.Attachments = nil
	}

	return s.newsletterRepository.UpdateNewsletter(*existingNewsletter)
}

func (s *NewsletterService) DeleteNewsletter(id string) error {
	return s.newsletterRepository.DeleteNewsletterByID(id)
}
