package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	domain "newsletter-app/pkg/domain/models"
	"newsletter-app/pkg/domain/ports"
	"newsletter-app/pkg/infrastructure/adapters/email"
	"newsletter-app/pkg/service"
	"newsletter-app/pkg/service/Dtos/request"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary Send newsletter to subscribers
// @Description Allows an admin user to send a newsletter to a list of subscribers
// @Tags newsletters
// @Accept json
// @Produce json
// @Param newsletterID path string true "ID of the newsletter to be sent"
// @Success 200 {string} string "OK"
// @Failure 400 {object} service.ErrorResponse "Bad Request"
// @Failure 500 {object} service.ErrorResponse "Internal Server Error"
// @Router /newsletters/send/{newsletterID} [post]
func SendNewsletterHandler(subscriberService ports.SubscriberServicePort, newsletterService ports.NewsletterServicePort, emailSender email.EmailSender) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newsletterID := mux.Vars(r)["newsletterID"]
		if newsletterID == "" {
			service.RespondWithError(w, http.StatusBadRequest, "Invalid newsletter ID")
			return
		}

		err := newsletterService.SendNewsletter(w, r, newsletterID, emailSender)
		if err != nil {
			fmt.Printf("Error sending newsletter: %s\n", err.Error())
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to send newsletter")
			return
		}

		service.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
			"status":  "OK",
			"message": "Newsletter sent successfully",
		})
	}
}

// @Summary Create a new newsletter
// @Description Allows an admin user to create a new newsletter
// @Tags newsletters
// @Accept json
// @Produce json
// @Param newsletter body domain.Newsletter true "Newsletter details"
// @Success 201 {string} string "Created"
// @Failure 400 {object} service.ErrorResponse "Bad Request"
// @Failure 500 {object} service.ErrorResponse "Internal Server Error"
// @Router /newsletters [post]
func CreateNewsletterHandler(newsletterService ports.NewsletterServicePort) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newNewsletter domain.Newsletter
		err := json.NewDecoder(r.Body).Decode(&newNewsletter)
		if err != nil {
			service.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		if newNewsletter.Category == "" {
			service.RespondWithError(w, http.StatusBadRequest, "Category is required")
			return
		}

		existingNewsletter, err := newsletterService.GetNewsletterByCategory(newNewsletter.Category)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to check for existing newsletters")
			return
		}

		if existingNewsletter != nil {
			service.RespondWithError(w, http.StatusConflict, "Newsletter with the specified category already exists")
			return
		}

		err = newsletterService.SaveNewsletter(newNewsletter)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to create newsletter")
			return
		}

		service.RespondWithJSON(w, http.StatusCreated, map[string]interface{}{
			"status":  "OK",
			"message": "Newsletter created successfully",
		})
	}
}

// @Summary Get a list of newsletters
// @Description Retrieves a list of newsletters with optional search and pagination parameters
// @Tags newsletters
// @Accept json
// @Produce json
// @Param name query string false "Name of the newsletter to search for"
// @Param page query int false "Page number for pagination"
// @Param pageSize query int false "Number of items per page for pagination"
// @Failure 400 {object} service.ErrorResponse "Bad Request"
// @Failure 500 {object} service.ErrorResponse "Internal Server Error"
// @Router /newsletters [get]
func GetNewslettersHandler(newsletterService ports.NewsletterServicePort) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

		newsletters, err := newsletterService.GetNewsletters(name, page, pageSize)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve newsletters")
			return
		}

		service.RespondWithJSON(w, http.StatusOK, newsletters)
	}
}

// @Summary Update an existing newsletter
// @Description Allows an admin user to update an existing newsletter
// @Tags newsletters
// @Accept json
// @Produce json
// @Param updateRequest body request.UpdateNewsletterRequest true "Update newsletter details"
// @Success 200 {string} string "OK"
// @Failure 400 {object} service.ErrorResponse "Bad Request"
// @Failure 500 {object} service.ErrorResponse "Internal Server Error"
// @Router /newsletters [put]
func UpdateNewsletterHandler(newsletterService ports.NewsletterServicePort) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var updateRequest request.UpdateNewsletterRequest
		err := json.NewDecoder(r.Body).Decode(&updateRequest)
		if err != nil {
			service.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}

		err = newsletterService.UpdateNewsletter(updateRequest)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to update newsletter")
			return
		}

		service.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
			"status":  "OK",
			"message": "Newsletter updated successfully",
		})
	}
}

// @Summary Delete a newsletter
// @Description Allows an admin user to delete a newsletter
// @Tags newsletters
// @Accept json
// @Produce json
// @Param id path string true "ID of the newsletter to delete"
// @Success 200 {string} string "OK"
// @Failure 400 {object} service.ErrorResponse "Bad Request"
// @Failure 500 {object} service.ErrorResponse "Internal Server Error"
// @Router /newsletters/{id} [delete]
func DeleteNewsletterHandler(newsletterService ports.NewsletterServicePort) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		if id == "" {
			service.RespondWithError(w, http.StatusBadRequest, "Invalid or missing newsletter ID")
			return
		}

		err := newsletterService.DeleteNewsletter(id)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to delete newsletter")
			return
		}

		service.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
			"status":  "OK",
			"message": "Newsletter deleted successfully",
		})
	}
}
