package model

import "time"

// Model :
type Model struct {
	CreatedAt     time.Time `bson:"createdAt"`
	UpdatedAt     time.Time `bson:"updatedAt"`
	IsSoftDeleted bool      `bson:"isSoftDeleted"`
}

// Collection name :
const (
	CollectionUser        = "User"
	CollectionFeature     = "Feature"
	CollectionUserFeature = "UserFeature"
)
