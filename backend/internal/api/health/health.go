package health

import (
	"context"
	"backend/internal/generated/api"
	"strings"
)

type Health struct{}

func New() *Health {
	return &Health{}
}

func (h *Health) HealthGet(ctx context.Context) (api.HealthGetOK, error) {
	return api.HealthGetOK{
		Data: strings.NewReader("OK"),
	}, nil
}
