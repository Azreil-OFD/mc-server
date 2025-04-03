package delivery_usecase

import (
	"context"
	"io"
	"os"
)

func (d *DownloadUsecase) GetModsArchive(ctx context.Context) (io.Reader, int64, error) {
	file, err := os.Open(d.archivePath)
	if err != nil {
		return nil, 0, err
	}
	stat, _ := file.Stat()
	return file, stat.Size(), nil
}
