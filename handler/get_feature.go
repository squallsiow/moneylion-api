package handler

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/moneylion-api/dto"
)

// GetFeatureAccess :
func (h *Handler) GetFeatureAccess(c echo.Context) error {

	var i struct {
		FeatureName string `json:"featureName" query:"featureName" validate:"required"`
		Email       string `json:"email"  query:"email" validate:"required"`
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

	// Find existing record
	userFeature, err := h.Repository.FindUserFeatureByFeatureNameUserEmail(ctx, i.FeatureName, i.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.NewError(http.StatusBadRequest, "feature access not found"))
	}

	return c.JSON(http.StatusOK, struct {
		CanAccess bool `json:"canAccess"`
	}{
		CanAccess: userFeature.Enable,
	})
}
