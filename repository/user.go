package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/moneylion-api/model"
)

// FindUserByEmail:
func (r *Repository) FindUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	err = r.db.Collection(model.CollectionUser).
		FindOne(ctx, bson.D{{"email", email}}).
		Decode(&user)
	if err != nil {
		return
	}
	return
}

func (r *Repository) CreateUser(ctx context.Context, user *model.User) (err error) {
	if user == nil {
		err = errors.New("user is nil")
		return
	}

	_, err = r.db.Collection(model.CollectionUser).
		InsertOne(ctx, user)
	if err != nil {
		return
	}
	return
}

// UpdateUserByID :
func (r *Repository) UpdateUserByID(ctx context.Context, id primitive.ObjectID, update bson.M) (user model.User, err error) {

	err = r.db.Collection(model.CollectionUser).
		FindOneAndUpdate(ctx, bson.D{{"_id", id}}, update,
			options.FindOneAndUpdate().
				SetReturnDocument(options.After)).
		Decode(&user)
	if err != nil {
		return
	}
	return
}

// UpdateUserByEmail :
func (r *Repository) UpdateUserByEmail(ctx context.Context, email string, update bson.M) (user model.User, err error) {

	err = r.db.Collection(model.CollectionUser).
		FindOneAndUpdate(ctx, bson.D{{"email", email}}, update,
			options.FindOneAndUpdate().
				SetReturnDocument(options.After)).
		Decode(&user)
	if err != nil {
		return
	}
	return
}

// DeleteUserByID :
func (r *Repository) DeleteUserByID(ctx context.Context, id string) (err error) {
	_, err = r.db.Collection(model.CollectionUser).
		DeleteOne(ctx, bson.D{{"_id", id}})

	if err != nil {
		return
	}
	return
}
