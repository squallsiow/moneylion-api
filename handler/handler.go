package handler

import "go.mongodb.org/mongo-driver/mongo"

// Handler : Singleton Handler
type Handler struct {
	client *mongo.Client
}

// New :
func New() (*Handler, error) {
	ctrl := &Handler{}
	return ctrl, nil
}
