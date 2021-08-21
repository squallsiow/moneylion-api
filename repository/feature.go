package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/moneylion-api/model"
)

// FindFeatureByName:
func (r *Repository) FindFeatureByName(ctx context.Context, name string) (feature model.Feature, err error) {
	err = r.db.Collection(model.CollectionFeature).
		FindOne(ctx, bson.D{{"name", name}}).
		Decode(&feature)
	if err != nil {
		return
	}
	return
}

func (r *Repository) CreateFeature(ctx context.Context, feature *model.Feature) (err error) {
	if feature == nil {
		err = errors.New("feature is nil")
		return
	}

	_, err = r.db.Collection(model.CollectionFeature).
		InsertOne(ctx, feature)
	if err != nil {
		return
	}
	return
}

// UpdateFeatureByID :
func (r *Repository) UpdateFeatureByID(ctx context.Context, id primitive.ObjectID, update bson.M) (feature model.Feature, err error) {

	err = r.db.Collection(model.CollectionFeature).
		FindOneAndUpdate(ctx, bson.D{{"_id", id}}, update,
			options.FindOneAndUpdate().
				SetReturnDocument(options.After)).
		Decode(&feature)
	if err != nil {
		return
	}
	return
}

// DeleteFeatureByID :
func (r *Repository) DeleteFeatureByID(ctx context.Context, id string) (err error) {
	_, err = r.db.Collection(model.CollectionFeature).
		DeleteOne(ctx, bson.D{{"_id", id}})

	if err != nil {
		return
	}
	return
}
