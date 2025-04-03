package delivery

import (
	"context"
	"io"
	"os"
	"backend/internal/generated/api"
)

type Delivery struct {
	usecase DownloadUsecase
}

type DownloadUsecase interface {
	GetModsArchive(ctx context.Context) (io.Reader, int64, error)
}

func NewDownloadHandler(usecase DownloadUsecase) *Delivery {
	return &Delivery{usecase: usecase}
}

func (d *Delivery) DownloadGet(ctx context.Context) (api.DownloadGetRes, error) {
	archive, size, err := d.usecase.GetModsArchive(ctx)
	if err != nil {
		if os.IsNotExist(err) {
			return &api.DownloadGetNotFound{}, nil
		}
		return &api.DownloadGetInternalServerError{}, nil
	}
	return &api.DownloadGetOKHeaders{
		ContentLength: api.NewOptInt64(size),
		Response:      api.DownloadGetOK{Data: archive},
	}, nil
}
