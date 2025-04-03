package upload

import (
	"context"
	"backend/internal/generated/api"
	"backend/internal/infrastructure/entity"

	ht "github.com/ogen-go/ogen/http"
	"github.com/samber/lo"
)

type Upload struct {
	u uploadUsecase
}

type uploadUsecase interface {
	UploadFiles(ctx context.Context, files []entity.File) (entity.UploadResult, error)
}

func New(u uploadUsecase) *Upload {
	return &Upload{
		u: u,
	}
}

func (h *Upload) UploadPost(ctx context.Context, req *api.UploadPostReq) (api.UploadPostRes, error) {
	files := lo.Map(req.Files, func(fh ht.MultipartFile, _ int) entity.File {
		return entity.File{
			Name: fh.Name,
			Size: fh.Size,
			File: fh.File,
		}
	})

	result, err := h.u.UploadFiles(ctx, files)
	if err != nil {
		return &api.UploadPostInternalServerError{}, nil
	}

	conflicts := lo.FilterMap(result.Skipped, func(s entity.UploadSkip, _ int) (string, bool) {
		return s.Name, s.Reason == "conflict"
	})
	invalid := lo.FilterMap(result.Skipped, func(s entity.UploadSkip, _ int) (string, bool) {
		return s.Name, s.Reason == "invalid"
	})

	uploaded := result.Uploaded
	skipped := len(result.Skipped)

	if len(invalid) == 0 {
		invalid = nil
	}	
	if len(conflicts) == 0 {
		conflicts = nil
	}

	switch {
	case len(uploaded) == 0 && skipped > 0:
		return &api.UploadPostConflict{
			Message:   api.NewOptString("Ни один файл не был загружен"),
			Conflicts: conflicts,
			Invalid:   invalid,
		}, nil

	case len(uploaded) > 0 && skipped > 0:
		return &api.UploadPostMultiStatus{
			Message:   api.NewOptString("Некоторые файлы не были загружены"),
			Uploaded:  uploaded,
			Conflicts: conflicts,
			Invalid:   invalid,
		}, nil

	default:
		return &api.UploadSuccessResponse{
			Message:  api.NewOptString("Все файлы успешно загружены"),
			Uploaded: uploaded,
		}, nil
	}
}
