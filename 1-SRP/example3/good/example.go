package goodsrp

// FileHandler is a struct that handles file operations.
type FileReader interface {
	Read(filepath string) ([]byte, error)
}

// FileValidator is an interface that validates file formats.
type FileValidator interface {
	ValidateFormat(content []byte) bool
}

// FileCompressor is an interface that compresses file content.
type FileCompressor interface {
	Compress(content []byte) ([]byte, error)
}

// CloudUploader is an interface that uploads files to cloud storage.
type CloudUploader interface {
	Upload(content []byte, destination string) error
}

type FileProcessorService struct {
	fileReader     FileReader
	fileValidator  FileValidator
	fileCompressor FileCompressor
	cloudUploader  CloudUploader
}

// NewFileProcessorService creates a new FileProcessorService with the provided dependencies.
func NewFileProcessorService(
	fileReader FileReader,
	fileValidator FileValidator,
	fileCompressor FileCompressor,
	cloudUploader CloudUploader,
) *FileProcessorService {
	return &FileProcessorService{
		fileReader:     fileReader,
		fileValidator:  fileValidator,
		fileCompressor: fileCompressor,
		cloudUploader:  cloudUploader,
	}
}

// ProcessFile reads, validates, compresses, and uploads a file.
func (s *FileProcessorService) ProcessFile(filepath string, destination string) error {
	content, err := s.fileReader.Read(filepath)
	if err != nil {
		return err
	}
	if !s.fileValidator.ValidateFormat(content) {
		return err
	}
	compressedContent, err := s.fileCompressor.Compress(content)
	if err != nil {
		return err
	}
	if err := s.cloudUploader.Upload(compressedContent, destination); err != nil {
		return err
	}
	return nil
}
