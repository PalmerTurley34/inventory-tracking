package models

import (
	"time"

	"github.com/PalmerTurley34/inventory-tracking/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Username  string
	ApiKey    string
}

func DBUserToResponse(user database.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Username:  user.Username,
		ApiKey:    user.ApiKey,
	}
}
