package handlers

import (
	"fmt"
	"net/http"
	"newsletter-app/pkg/domain/ports"
	"newsletter-app/pkg/service"
	"strconv"

	"github.com/gorilla/mux"
)

// @Summary Subscribe to the newsletter
// @Description Allows a user to subscribe to the newsletter
// @Tags subscribers
// @Accept json
// @Produce json
// @Param email path string true "Email address to subscribe"
// @Param category path string true "Category to subscribe to"
// @Success 200 {string} string "OK"
// @Failure 400 {object} service.ErrorResponse "Bad Request"
// @Failure 500 {object} service.ErrorResponse "Internal Server Error"
// @Router /subscribe/{email}/{category} [post]
func SubscribeHandler(subscriberService ports.SubscriberServicePort) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := mux.Vars(r)["email"]
		category := mux.Vars(r)["category"]
		if email == "" || !service.IsValidEmail(email) {
			service.RespondWithError(w, http.StatusBadRequest, "Invalid or missing email address")
			return
		}

		existingSubscriber, err := subscriberService.GetSubscriberByEmail(email, category)
		if err == nil && existingSubscriber != nil {
			fmt.Println("Email is invalid or missing:", email)
			service.RespondWithError(w, http.StatusConflict, "User is already subscribed")
			return
		}

		err = subscriberService.Subscribe(email, category)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to subscribe user")
			return
		}

		subscriber, err := subscriberService.GetSubscriberByEmail(email, category)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to get subscriber details")
			return
		}

		service.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
			"status":     "OK",
			"message":    "Subscription successful",
			"subscriber": subscriber,
		})
	}
}

// @Summary Unsubscribe from the newsletter
// @Description Allows a user to unsubscribe from the newsletter
// @Tags subscribers
// @Accept json
// @Produce json
// @Param email path string true "Email address to unsubscribe"
// @Param category path string true "Category to subscribe to"
// @Success 200 {string} string "OK"
// @Failure 400 {object} service.ErrorResponse "Bad Request"
// @Failure 500 {object} service.ErrorResponse "Internal Server Error"
// @Router /unsubscribe/{email}/{category} [delete]
func UnsubscribeHandler(subscriberService ports.SubscriberServicePort) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := mux.Vars(r)["email"]
		category := mux.Vars(r)["category"]
		if email == "" || !service.IsValidEmail(email) {
			service.RespondWithError(w, http.StatusBadRequest, "Invalid or missing email address")
			return
		}

		err := subscriberService.Unsubscribe(email, category)
		if err != nil {
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to unsubscribe user")
			return
		}

		service.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
			"status":  "OK",
			"message": "User unsubscribed successfully",
		})
	}
}

// @Summary Get subscriber by email and category
// @Description Get details of a subscriber by email address
// @Tags subscribers
// @Accept json
// @Produce json
// @Param email path string true "Email address to get details"
// @Param category path string true "Category to subscribe to"
// @Success 200 {object} domain.Subscriber
// @Failure 404 {object} service.ErrorResponse "Subscriber not found"
// @Failure 500 {object} service.ErrorResponse "Internal Server Error"
// @Router /subscribers/{email}/{category} [get]
func GetSubscriberHandler(subscriberService ports.SubscriberServicePort) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := mux.Vars(r)["email"]
		category := mux.Vars(r)["category"]
		if email == "" || !service.IsValidEmail(email) {
			service.RespondWithError(w, http.StatusBadRequest, "Invalid or missing email address")
			return
		}

		subscriber, err := subscriberService.GetSubscriberByEmail(email, category)
		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				service.RespondWithError(w, http.StatusNotFound, "Subscriber not found")
				return
			}

			service.RespondWithError(w, http.StatusInternalServerError, "Failed to get subscriber")
			return
		}

		service.RespondWithJSON(w, http.StatusOK, subscriber)
	}
}

// @Summary Get a list of subscribers
// @Description Retrieves a list of subscribers with optional search and pagination parameters
// @Tags subscribers
// @Accept json
// @Produce json
// @Param email query string false "Email address of the subscriber to search for"
// @Param category query string false "Category of the subscriber to search for"
// @Param page query int false "Page number for pagination"
// @Param pageSize query int false "Number of items per page for pagination"
// @Failure 400 {object} service.ErrorResponse "Bad Request"
// @Failure 500 {object} service.ErrorResponse "Internal Server Error"
// @Router /subscribers [get]
func GetSubscribersHandler(subscriberService ports.SubscriberServicePort) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		category := r.URL.Query().Get("category")
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

		subscribers, err := subscriberService.GetSubscribers(email, category, page, pageSize)
		if err != nil {
			fmt.Println("Email is invalid or missing:", email)
			service.RespondWithError(w, http.StatusInternalServerError, "Failed to retrieve subscribers")
			return
		}
		service.RespondWithJSON(w, http.StatusOK, subscribers)
	}
}
