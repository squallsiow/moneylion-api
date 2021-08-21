package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	client *mongo.Client
	db     *mongo.Database
}

// New :
func New(dbName string, mongo *mongo.Client) *Repository {
	return &Repository{
		client: mongo,
		db:     mongo.Database(dbName),
	}
}
