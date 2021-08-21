package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// UserFeature:User and Feature MN object
type UserFeature struct {
	UserID    primitive.ObjectID `bson:"userID"`
	FeatureID primitive.ObjectID `bson:"featureID"`
	Enable    bool               `bson:"enable"`
	Model     `bson:",inline"`
}
