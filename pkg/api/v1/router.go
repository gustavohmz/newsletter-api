package v1

import (
	"newsletter-app/pkg/api/v1/handlers"
	"newsletter-app/pkg/domain/ports"
	"newsletter-app/pkg/infrastructure/adapters/email"
	"newsletter-app/pkg/infrastructure/adapters/mongodb"
	"newsletter-app/pkg/service"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	subscriberRepo := mongodb.NewSubscriberRepository()
	newsletterRepo := mongodb.NewNewsletterRepository()

	var subscriberService ports.SubscriberServicePort = service.NewSubscriberService(subscriberRepo)
	var newsletterService ports.NewsletterServicePort = service.NewNewsletterService(newsletterRepo, subscriberRepo)

	var emailSender email.EmailSender = email.NewBrevoEmailSender()

	// Routes configuration for subscribers
	r.HandleFunc("/api/v1/subscribe/{email}/{category}", handlers.SubscribeHandler(subscriberService)).Methods("POST")
	r.HandleFunc("/api/v1/unsubscribe/{email}/{category}", handlers.UnsubscribeHandler(subscriberService)).Methods("DELETE")
	r.HandleFunc("/api/v1/subscribers/{email}/{category}", handlers.GetSubscriberHandler(subscriberService)).Methods("GET")
	r.HandleFunc("/api/v1/subscribers", handlers.GetSubscribersHandler(subscriberService)).Methods("GET")

	// Routes configuration for newsletters
	r.HandleFunc("/api/v1/newsletters/send/{newsletterID}", handlers.SendNewsletterHandler(subscriberService, newsletterService, emailSender)).Methods("POST")
	r.HandleFunc("/api/v1/newsletters", handlers.CreateNewsletterHandler(newsletterService)).Methods("POST")
	r.HandleFunc("/api/v1/newsletters", handlers.GetNewslettersHandler(newsletterService)).Methods("GET")
	r.HandleFunc("/api/v1/newsletters", handlers.UpdateNewsletterHandler(newsletterService)).Methods("PUT")
	r.HandleFunc("/api/v1/newsletters/{id}", handlers.DeleteNewsletterHandler(newsletterService)).Methods("DELETE")

	return r
}
