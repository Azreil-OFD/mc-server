package delivery_usecase

type DownloadUsecase struct {
	archivePath string
}

func New(archivePath string) *DownloadUsecase {
	return &DownloadUsecase{archivePath: archivePath}
}
