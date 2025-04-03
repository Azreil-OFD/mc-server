package list_usecase

import (
	"context"
	"os"
	"backend/internal/infrastructure/entity"
	"strings"
)

type ListUsecase struct {
}

func New() *ListUsecase {
	return &ListUsecase{}
}

func (l *ListUsecase) ListGet(ctx context.Context) ([]entity.List, error) {
	files, err := os.ReadDir("../mods")
	if err != nil {
		return nil, err
	}

	var result []entity.List
	for _, f := range files {
		result = append(result, entity.List{
			Name:   f.Name(),
			Active: isMod(f.Name()),
		})
	}

	return result, nil
}

// TODO проверить
func isMod(fileName string) bool {
	// Проверяем, оканчивается ли имя строго на ".jar"
	return strings.HasSuffix(fileName, ".jar") &&
		len(fileName) == len(strings.TrimSuffix(fileName, ".jar"))+4
}
