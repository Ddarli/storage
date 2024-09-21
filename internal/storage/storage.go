package storage

import (
	"fmt"
	"github.com/google/uuid"
	"storage/internal/file"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{files: make(map[uuid.UUID]*file.File)}
}
func (s *Storage) SaveFile(filename string, data []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, data)
	if err != nil {
		return nil, err
	}
	s.files[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetByID(uuid uuid.UUID) (*file.File, error) {
	newFile, ok := s.files[uuid]
	if !ok {
		return nil, fmt.Errorf("newFile %v not found", uuid)
	}
	return newFile, nil
}
