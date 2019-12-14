package model

import "time"

// Model :
type Model struct {
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	IsSoftDeleted bool      `json:"isSoftDeleted"`
}
