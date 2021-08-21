package dto

type UserFeature struct {
	FeatureName string `json:"featureName"`
	Email       string `json:"email"`
	Enable      bool   `json:"enable"`
}
