package transformer

import (
	"github.com/moneylion-api/dto"
	"github.com/moneylion-api/model"
)

// ToUserFeature:
func ToUserFeature(user model.User, feature model.Feature, enable bool) dto.UserFeature {
	obj := dto.UserFeature{
		FeatureName: feature.Name,
		Email:       user.Email,
		Enable:      enable,
	}
	return obj
}
