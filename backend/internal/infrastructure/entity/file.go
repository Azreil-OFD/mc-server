package entity

import "io"

type File struct {
	Name string
	Size int64
	File io.Reader
}

type UploadSkip struct {
	Name   string // имя файла
	Reason string // "conflict" или "invalid"
}

type UploadResult struct {
	Uploaded []string       // успешно загруженные файлы
	Skipped  []UploadSkip   // пропущенные с причинами
}

