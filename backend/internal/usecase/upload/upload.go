package upload_usecase

import (
	"context"
	"errors"
	"io"
	"os"
	"path/filepath"
	"backend/internal/infrastructure/entity"
	"strings"
)

type UploadUsecase struct {
}

func New() *UploadUsecase {
	return &UploadUsecase{}
}

var dirMods = "mods"

func (u *UploadUsecase) UploadFiles(ctx context.Context, files []entity.File) (entity.UploadResult, error) {
	var result entity.UploadResult

	for _, f := range files {
		if !strings.HasSuffix(f.Name, ".jar") {
			result.Skipped = append(result.Skipped, entity.UploadSkip{
				Name:   f.Name,
				Reason: "invalid",
			})
			continue
		}

		dstPath := filepath.Join(dirMods, filepath.Base(f.Name))

		if _, statErr := os.Stat(dstPath); statErr == nil {
			result.Skipped = append(result.Skipped, entity.UploadSkip{
				Name:   f.Name,
				Reason: "conflict",
			})
			continue
		} else if !errors.Is(statErr, os.ErrNotExist) {
			return entity.UploadResult{}, statErr
		}

		dst, err := os.Create(dstPath)
		if err != nil {
			u.rollbackFiles(result.Uploaded)
			return entity.UploadResult{}, err
		}
		defer dst.Close()

		if _, err := io.Copy(dst, f.File); err != nil {
			u.rollbackFiles(result.Uploaded)
			return entity.UploadResult{}, err
		}

		if closer, ok := f.File.(io.Closer); ok {
			_ = closer.Close()
		}

		result.Uploaded = append(result.Uploaded, f.Name)
	}

	return result, nil
}

func (u *UploadUsecase) rollbackFiles(paths []string) {
	for _, path := range paths {
		_ = os.Remove(filepath.Join(dirMods, filepath.Base(path)))
	}
}
