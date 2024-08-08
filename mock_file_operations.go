// mock_file_operations.go
package main

import (
	"github.com/stretchr/testify/mock"
)

// MockFileOperations is a mock implementation of the FileOperations interface.
type MockFileOperations struct {
	mock.Mock
}

func (m *MockFileOperations) createFile(username, foldername, filename, description string) error {
	args := m.Called(username, foldername, filename, description)
	return args.Error(0)
}

func (m *MockFileOperations) deleteFile(username, foldername, filename string) error {
	args := m.Called(username, foldername, filename)
	return args.Error(0)
}
