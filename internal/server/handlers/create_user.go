package handlers

import (
	"net/http"

	"github.com/SKilliu/cardgame/internal/db"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (h *Handler) CreateUser(c echo.Context) error {
	var req CreateUserRequest

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse create order request")
		return c.JSON(http.StatusBadRequest, "bad param in body")
	}

	tx, err := h.db.NewSession(nil).Begin()
	if err != nil {
		h.log.WithError(err).Error("failed to create db transaction")
		return c.JSON(http.StatusBadRequest, "bad param in body")
	}
	defer tx.RollbackUnlessCommitted()

	_, err = db.Insert(tx, db.User{
		ID:             uuid.New().String(),
		Name:           req.Name,
		HashedPassword: req.Password,
		Email:          req.Email,
	})
	if err != nil {
		h.log.WithError(err).Error("failed to create new user")
		return c.JSON(http.StatusInternalServerError, "failed to create new user")
	}

	err = tx.Commit()
	if err != nil {
		h.log.WithError(err).Error("failed to commit a transaction")
		return c.JSON(http.StatusInternalServerError, "failed to commit a transaction")
	}

	return c.NoContent(http.StatusOK)
}
