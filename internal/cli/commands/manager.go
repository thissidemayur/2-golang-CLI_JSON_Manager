package commands

import (
	"errors"
	"fmt"
)

// sentimental error (resuable constant)
var (
	ErrNotFound = errors.New("record not found")
	ErrInvalidId = errors.New("invalid ID")
	ErrFileNotFound = errors.New("file not found")
	ErrEmptyName = errors.New("name cannot be empty")
)
// Manager struct to handle records
type Manager struct {
	fileName string
}

// NewManager creates a new Manager instance(constructor)
func NewManager(fileName string) (*Manager ,error ){
	m:= &Manager{fileName: fileName}
	if err := m.EnsureFile(); err != nil {
		return nil, fmt.Errorf("ensure file failed: %w", err)
	}
	return m, nil
}

