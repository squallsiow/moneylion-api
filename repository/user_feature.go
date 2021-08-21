package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/moneylion-api/model"
	"go.mongodb.org/mongo-driver/bson"
)

// FindUserFeatureByFeatureNameUserName:
func (r *Repository) FindUserFeatureByFeatureNameUserEmail(ctx context.Context, featureName string, email string) (userFeature *model.UserFeature, err error) {
	session, err := r.client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	if err := session.StartTransaction(); err != nil {
		return nil, err
	}

	if err := mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		var feature model.Feature
		err = r.db.Collection(model.CollectionFeature).
			FindOne(sc, bson.D{{"name", featureName}}).
			Decode(&feature)
		if err != nil {
			return err
		}
		var user model.User
		err = r.db.Collection(model.CollectionUser).
			FindOne(sc, bson.D{{"email", email}}).
			Decode(&user)
		if err != nil {
			return err
		}

		var uf model.UserFeature
		err = r.db.Collection(model.CollectionUserFeature).
			FindOne(sc, bson.M{
				"featureID": feature.ID,
				"userID":    user.ID,
			}).Decode(&uf)
		if err != nil {
			return err
		}

		userFeature = &uf

		return nil
	}); err != nil {
		defer session.AbortTransaction(ctx)
		return nil, err
	}

	return userFeature, nil
}

func (r *Repository) CreateUserFeature(ctx context.Context, userfeature *model.UserFeature) (err error) {
	if userfeature == nil {
		err = errors.New("userfeature is nil")
		return
	}

	_, err = r.db.Collection(model.CollectionUserFeature).
		InsertOne(ctx, userfeature)
	if err != nil {
		return
	}
	return
}

// DeleteUserFeatureByUserID :
func (r *Repository) DeleteUserFeatureByUserID(ctx context.Context, id string) (err error) {
	_, err = r.db.Collection(model.CollectionUserFeature).
		DeleteOne(ctx, bson.D{{"userID", id}})

	if err != nil {
		return
	}
	return
}

// UpdateUserFeatureByIDs :
func (r *Repository) UpdateUserFeatureByIDs(
	ctx context.Context,
	userID primitive.ObjectID,
	featureID primitive.ObjectID,
	update bson.M) (userFeature model.UserFeature, err error) {

	err = r.db.Collection(model.CollectionUserFeature).
		FindOneAndUpdate(
			ctx,
			bson.M{"userID": userID,
				"featureID": featureID},
			update,
			options.FindOneAndUpdate().
				SetReturnDocument(options.After)).
		Decode(&userFeature)
	if err != nil {
		return
	}
	return
}
