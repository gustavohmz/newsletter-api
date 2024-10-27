package request

import "go.mongodb.org/mongo-driver/bson/primitive"

// UpdateNewsletterRequest represents the structure for the newsletter update request.
type UpdateNewsletterRequest struct {
	ID          primitive.ObjectID `json:"id"`
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
