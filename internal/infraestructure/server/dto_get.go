package server

import (
	"staffing-app-backend/internal/core/domain"
)

type ResponsePersonGet domain.Category

func BuildResponsePersonGet(category domain.Category) ResponsePersonGet {
	return ResponsePersonGet(category)
}
