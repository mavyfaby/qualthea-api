package category

import (
	db "qualthea-api/internal/db/models/category/db"
)

// Service define a service
type Service struct {
	r *db.Queries
}

// NewService creates a new service for category
func NewService(r *db.Queries) *Service {
	return &Service{r: r}
}
