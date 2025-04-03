package list

import (
	"context"
	"backend/internal/generated/api"
	"backend/internal/infrastructure/entity"

	"github.com/AlekSi/pointer"
	"github.com/samber/lo"
)

type List struct {
	list listUsecase
}

type listUsecase interface {
	ListGet(ctx context.Context) ([]entity.List, error)
}

func New(list listUsecase) *List {
	return &List{list: list}
}

func (l *List) ListGet(ctx context.Context) (api.ListGetRes, error) {
	entityList, err := l.list.ListGet(ctx)
	if err != nil {
		return &api.ListGetInternalServerError{}, nil
	}

	modInfo := lo.Map(entityList, func(item entity.List, _ int) api.ModInfo {
		return api.ModInfo{
			Name:   api.NewOptString(item.Name),
			Active: api.NewOptBool(item.Active),
		}
	})

	return pointer.To[api.ListGetOKApplicationJSON](modInfo), nil
}
