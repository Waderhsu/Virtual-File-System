package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileOperations_CreateFile(t *testing.T) {
	mockFileOps := new(MockFileOperations)
	mockFileOps.On("createFile", "user1", "folder1", "file.txt", "File description").Return(nil)

	err := mockFileOps.createFile("user1", "folder1", "file.txt", "File description")
	require.NoError(t, err)

	mockFileOps.AssertExpectations(t)
}

func TestCreateFileError(t *testing.T) {
	mockFileOps := new(MockFileOperations)
	mockFileOps.On("createFile", "user1", "folder1", "file.txt", "File description").Return(errors.New("file creation error"))

	err := mockFileOps.createFile("user1", "folder1", "file.txt", "File description")
	require.Error(t, err)
	require.Equal(t, "file creation error", err.Error())

	mockFileOps.AssertExpectations(t)
}

func TestFileOperations_DeleteFile(t *testing.T) {
	mockFileOps := new(MockFileOperations)
	mockFileOps.On("deleteFile", "user1", "folder1", "file.txt").Return(nil)

	err := mockFileOps.deleteFile("user1", "folder1", "file.txt")
	require.NoError(t, err)

	mockFileOps.AssertExpectations(t)
}

func TestDeleteFileError(t *testing.T) {
	mockFileOps := new(MockFileOperations)
	mockFileOps.On("deleteFile", "user1", "folder1", "file.txt").Return(errors.New("file deletion error"))

	err := mockFileOps.deleteFile("user1", "folder1", "file.txt")
	require.Error(t, err)
	require.Equal(t, "file deletion error", err.Error())

	mockFileOps.AssertExpectations(t)
}