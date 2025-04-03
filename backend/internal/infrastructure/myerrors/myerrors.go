package myerrors

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidExtension = errors.New("только .jar файлы разрешены")

type ConflictError struct {
	Files []string
}

func (e *ConflictError) Error() string {
	return fmt.Sprintf("Файлы уже существуют: %s", strings.Join(e.Files, ", "))
}

func NewConflictError(files []string) error {
	return &ConflictError{Files: files}
}
