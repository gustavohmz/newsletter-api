package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// represents a newsletter subscriber.
// swagger:model
type Subscriber struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email            string             `json:"email"`
	SubscriptionDate time.Time          `json:"subscription_date"`
	Category         string             `json:"category"`
}

type Subscribers []Subscriber
