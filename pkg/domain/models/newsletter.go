package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// represents a newsletter.
// swagger:model
type Newsletter struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" example:""`
	Name        string             `json:"name"`
	Category    string             `json:"category"`
	Subject     string             `json:"subject"`
	Content     string             `json:"content"`
	Attachments []Attachment       `json:"attachments"`
}

// represents a file attached to the newsletter.
// swagger:model
type Attachment struct {
	Name string `json:"name"`
	Data string `json:"data"`
	Type string `json:"type"`
}
