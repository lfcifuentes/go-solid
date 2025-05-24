package badsrp

import (
	"fmt"
)

// FileHandler is a struct that handles file operations.
type FileHandler struct {
	filepath string
	content  []byte
}

// ReadFile reads the content of a file.
func (f *FileHandler) ReadFile() error {
	fmt.Println("Reading file:", f.filepath)
	return nil
}

// ValidateFormat checks if the file format is valid.
// It should only validate the file format, not handle other operations.
// This violates the Single Responsibility Principle (SRP) because it mixes file reading and validation.
func (f *FileHandler) ValidateFormat() error {
	fmt.Println("Validating file format")
	return nil
}

// CompressFile compresses the file content.
// This method is responsible for compressing the file content, which is a separate concern from reading and validating the file.
// This violates the Single Responsibility Principle (SRP) because it combines multiple responsibilities in one struct.
// Ideally, file compression should be handled by a separate component or service.
func (f *FileHandler) CompressFile() error {
	fmt.Println("Compressing file")
	return nil
}

// UploadToCloud uploads the file to a cloud storage.
// This method is responsible for both file handling and cloud operations,
// which violates the Single Responsibility Principle (SRP).
// Ideally, file handling and cloud operations should be separated into different components.
func (f *FileHandler) UploadToCloud() error {
	fmt.Println("Uploading to cloud")
	return nil
}

// LogOperation logs the file operation.
// This method violates the Single Responsibility Principle
// because it mixes file handling with logging functionality.
// Ideally, logging should be handled by a separate component.
func (f *FileHandler) LogOperation() {
	fmt.Println("Logging operation")
}
