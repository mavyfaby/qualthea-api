package user

import (
	db "qualthea-api/internal/db/models/user/db"
)

// Service define a service
type Service struct {
	r *db.Queries
}

// NewService creates a new service for category
func NewService(r *db.Queries) *Service {
	return &Service{r: r}
}
