package api

import (
	"backend/internal/api/delivery"
	"backend/internal/api/health"
	"backend/internal/api/list"
	"backend/internal/api/upload"
)

type API struct {
	*delivery.Delivery
	*health.Health
	*list.List
	*upload.Upload
}
