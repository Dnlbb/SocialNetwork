package utils

import "social_network/models"

type AuthStorage struct {
    Users map[string]models.User
}
