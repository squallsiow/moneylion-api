package handler

import (
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/moneylion-api/model"

	"github.com/labstack/echo/v4"
	"github.com/moneylion-api/dto"
	"github.com/moneylion-api/transformer"
)

// AccessFeature :
func (h *Handler) AccessFeature(c echo.Context) error {

	var i struct {
		FeatureName string `json:"featureName" query:"featureName" validate:"required"`
		Email       string `json:"email"  query:"email" validate:"required"`
		Enable      bool   `json:"enable" query:"enable"`
	}

	// bind input value to variable i
	if err := c.Bind(&i); err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.NewError(http.StatusBadRequest, err.Error()),
		)
	}

	//validation input
	if err := c.Validate(i); err != nil {
		return c.JSON(http.StatusUnprocessableEntity,
			dto.NewError(http.StatusUnprocessableEntity, err.Error()))
	}

	// Sanitize value
	i.FeatureName = strings.TrimSpace(i.FeatureName)
	i.Email = strings.TrimSpace(i.Email)

	ctx := c.Request().Context()

	user, err := h.Repository.FindUserByEmail(ctx, i.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.NewError(http.StatusBadRequest, "user not exist"))
	}

	feature, err := h.Repository.FindFeatureByName(ctx, i.FeatureName)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			dto.NewError(http.StatusBadRequest, "feature not exist"))
	}

	// Find existing record
	userFeature, err := h.Repository.FindUserFeatureByFeatureNameUserEmail(ctx, i.FeatureName, i.Email)
	if err == nil {
		// if record exist, then need to determine whether it is update or no change
		if userFeature.Enable != i.Enable {
			_, err := h.Repository.UpdateUserFeatureByIDs(
				ctx,
				user.ID,
				feature.ID,
				bson.M{"$set": bson.M{
					"enable": i.Enable,
				}},
			)
			if err != nil {
				return c.JSON(http.StatusInternalServerError,
					dto.NewError(http.StatusInternalServerError, err.Error()))
			}
			return c.JSON(http.StatusOK, transformer.ToUserFeature(user, feature, i.Enable))
		}
		return c.JSON(http.StatusNotModified, transformer.ToUserFeature(user, feature, i.Enable))
	} else {
		//Create user feature
		userFeature := model.UserFeature{
			UserID:    user.ID,
			FeatureID: feature.ID,
			Enable:    i.Enable,
			Model: model.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}

		err := h.Repository.CreateUserFeature(ctx, &userFeature)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				dto.NewError(http.StatusInternalServerError, err.Error()))
		}
	}

	return c.JSON(http.StatusOK, transformer.ToUserFeature(user, feature, i.Enable))
}
