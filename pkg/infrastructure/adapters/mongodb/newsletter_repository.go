package mongodb

import (
	"context"
	domain "newsletter-app/pkg/domain/models"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NewsletterRepository struct {
	newsletterCollection *mongo.Collection
}

func NewNewsletterRepository() *NewsletterRepository {
	mongoDb := os.Getenv("mongoDb")
	mongoNewsletterCollection := os.Getenv("mongoNewsletterCollection")

	return &NewsletterRepository{
		newsletterCollection: client.Database(mongoDb).Collection(mongoNewsletterCollection),
	}
}

func (r *NewsletterRepository) SaveNewsletter(newsletter domain.Newsletter) error {
	_, err := r.newsletterCollection.InsertOne(context.TODO(), newsletter)
	return err
}

func (r *NewsletterRepository) GetNewsletterByID(newsletterID string) (*domain.Newsletter, error) {
	var newsletter domain.Newsletter
	objectID, err := primitive.ObjectIDFromHex(newsletterID)
	if err != nil {
		return nil, err
	}

	err = r.newsletterCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&newsletter)
	if err != nil {
		return nil, err
	}
	return &newsletter, nil
}

func (r *NewsletterRepository) UpdateNewsletter(newsletter domain.Newsletter) error {
	filter := bson.M{"_id": newsletter.ID}

	update := bson.M{"$set": bson.M{
		"name":        newsletter.Name,
		"category":    newsletter.Category,
		"subject":     newsletter.Subject,
		"content":     newsletter.Content,
		"attachments": newsletter.Attachments,
	}}

	_, err := r.newsletterCollection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r *NewsletterRepository) GetNewsletterByCategory(category string) (*domain.Newsletter, error) {
	var newsletter domain.Newsletter
	filter := bson.M{"category": category}

	err := r.newsletterCollection.FindOne(context.TODO(), filter).Decode(&newsletter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &newsletter, nil
}

func (r *NewsletterRepository) GetNewsletters(searchName string, page int, pageSize int) ([]domain.Newsletter, error) {
	filter := bson.M{}
	if searchName != "" {
		filter["name"] = primitive.Regex{Pattern: searchName, Options: "i"}
	}

	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * pageSize))
	findOptions.SetLimit(int64(pageSize))

	cursor, err := r.newsletterCollection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var newsletters []domain.Newsletter

	for cursor.Next(context.TODO()) {
		var newsletter domain.Newsletter
		if err := cursor.Decode(&newsletter); err != nil {
			return nil, err
		}
		newsletters = append(newsletters, newsletter)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return newsletters, nil
}

func (r *NewsletterRepository) DeleteNewsletterByID(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	_, err = r.newsletterCollection.DeleteOne(context.TODO(), filter)
	return err
}
