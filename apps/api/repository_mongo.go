package main

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongoTimeout = 5 * time.Second

type mongoContactRepository struct {
	collection *mongo.Collection
}

func NewMongoClient(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}

func NewMongoContactRepository(client *mongo.Client, dbName, collection string) ContactRepository {
	return &mongoContactRepository{
		collection: client.Database(dbName).Collection(collection),
	}
}

func (r *mongoContactRepository) Create(payload *contact) error {
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, payload)
	return err
}

func (r *mongoContactRepository) List() ([]contact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	cur, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	contacts := []contact{}
	if err := cur.All(ctx, &contacts); err != nil {
		return nil, err
	}
	return contacts, nil
}

func (r *mongoContactRepository) Get(id string) (*contact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	var co contact
	err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&co)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, ErrContactNotFound
	}
	if err != nil {
		return nil, err
	}
	return &co, nil
}

func (r *mongoContactRepository) Update(id string, payload *contact) error {
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"firstName":   payload.FirstName,
			"lastName":    payload.LastName,
			"phoneNumber": payload.PhoneNumber,
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"id": id}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return ErrContactNotFound
	}
	return nil
}

func (r *mongoContactRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()

	result, err := r.collection.DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return ErrContactNotFound
	}
	return nil
}
