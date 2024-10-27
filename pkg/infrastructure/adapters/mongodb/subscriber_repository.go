package mongodb

import (
	"context"
	domain "newsletter-app/pkg/domain/models"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubscriberRepository struct {
	subscriberCollection *mongo.Collection
}

func NewSubscriberRepository() *SubscriberRepository {
	mongoDb := os.Getenv("mongoDb")
	mongoSubscriberCollection := os.Getenv("mongoSubscriberCollection")

	return &SubscriberRepository{
		subscriberCollection: client.Database(mongoDb).Collection(mongoSubscriberCollection),
	}
}

func (r *SubscriberRepository) SaveSubscriber(subscriber domain.Subscriber) error {
	_, err := r.subscriberCollection.InsertOne(context.TODO(), subscriber)
	return err
}

func (r *SubscriberRepository) GetSubscriberByEmailAndCategory(email, category string) (*domain.Subscriber, error) {
	var subscriber domain.Subscriber
	filter := bson.M{"email": email, "category": category}
	err := r.subscriberCollection.FindOne(context.TODO(), filter).Decode(&subscriber)
	if err != nil {
		return nil, err
	}
	return &subscriber, nil
}

func (r *SubscriberRepository) GetSubscribers(email, category string, page, pageSize int) ([]domain.Subscriber, error) {
	var subscribers []domain.Subscriber

	filter := bson.M{}
	if email != "" {
		filter["email"] = email
	}
	if category != "" {
		filter["category"] = category
	}

	cursor, err := r.subscriberCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var subscriber domain.Subscriber
		if err := cursor.Decode(&subscriber); err != nil {
			return nil, err
		}
		subscribers = append(subscribers, subscriber)
	}

	return subscribers, nil
}

func (r *SubscriberRepository) DeleteSubscriberByEmail(email, category string) error {
	filter := bson.M{"email": email}

	if category != "" {
		filter["category"] = category
	}

	_, err := r.subscriberCollection.DeleteMany(context.TODO(), filter)
	return err
}

func (r *SubscriberRepository) GetSubscribersByCategory(category string) ([]domain.Subscriber, error) {
	filter := bson.M{"category": category}

	cursor, err := r.subscriberCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var subscribers []domain.Subscriber
	for cursor.Next(context.TODO()) {
		var subscriber domain.Subscriber
		if err := cursor.Decode(&subscriber); err != nil {
			return nil, err
		}
		subscribers = append(subscribers, subscriber)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return subscribers, nil
}
