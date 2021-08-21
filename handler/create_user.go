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

// CreateUser :
func (h *Handler) CreateUser(c echo.Context) error {

	var i struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required"`
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
	i.Email = strings.TrimSpace(i.Email)

	ctx := c.Request().Context()

	if _, err := h.Repository.FindUserByEmail(ctx, i.Email); err == nil {
		return c.JSON(http.StatusBadRequest,
			dto.NewError(http.StatusBadRequest, "email already exist"))
	}

	user := model.User{
		ID:    primitive.NewObjectID(),
		Name:  i.Name,
		Email: i.Email,
		Model: model.Model{
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			IsSoftDeleted: false,
		},
	}

	if err := h.Repository.CreateUser(ctx, &user); err != nil {
		return c.JSON(http.StatusInternalServerError,
			dto.NewError(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, transformer.ToUser(user))
}
