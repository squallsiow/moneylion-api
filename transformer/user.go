package transformer

import (
	"github.com/moneylion-api/dto"
	"github.com/moneylion-api/model"
)

// ToUser:
func ToUser(user model.User) dto.User {
	obj := dto.User{
		ID:    user.ID.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}
	return obj
}
