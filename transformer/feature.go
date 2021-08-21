package transformer

import (
	"github.com/moneylion-api/dto"
	"github.com/moneylion-api/model"
)

// ToFeature:
func ToFeature(feature model.Feature) dto.Feature {
	obj := dto.Feature{
		Name: feature.Name,
	}
	return obj
}
