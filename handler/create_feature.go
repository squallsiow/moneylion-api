package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/moneylion-api/model"
	"github.com/moneylion-api/transformer"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/moneylion-api/dto"

	"github.com/labstack/echo/v4"
)

// CreateFeature :
func (h *Handler) CreateFeature(c echo.Context) error {

	var i struct {
		Name string `json:"name" validate:"required"`
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
	i.Name = strings.TrimSpace(i.Name)

	ctx := c.Request().Context()

	if _, err := h.Repository.FindFeatureByName(ctx, i.Name); err == nil {
		return c.JSON(http.StatusBadRequest,
			dto.NewError(http.StatusBadRequest, "feature name already exist"))
	}

	feature := model.Feature{
		ID:   primitive.NewObjectID(),
		Name: i.Name,
		Model: model.Model{
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	}

	if err := h.Repository.CreateFeature(ctx, &feature); err != nil {
		return c.JSON(http.StatusInternalServerError,
			dto.NewError(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, transformer.ToFeature(feature))
}
