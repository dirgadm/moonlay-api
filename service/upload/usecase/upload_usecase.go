package usecase

import (
	"context"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"project-version3/moonlay-api/service/domain"
	"project-version3/moonlay-api/service/domain/dto"

	"github.com/labstack/gommon/log"
)

type uploadUsecase struct {
}

// NewUploadUsecase will create new an articleUsecase object representation of domain.UploadUsecase interface
func NewUploadUsecase(timeout time.Duration) domain.UploadUsecase {
	return &uploadUsecase{}
}

func (h uploadUsecase) UploadFile(ctx context.Context, w http.ResponseWriter, r *http.Request) (res []dto.UploadResponse, err error) {
	// Parse the multipart form data
	err = r.ParseMultipartForm(10 << 20) // 10 MB limit for file size
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Specify the allowed file extensions
	allowedExtensions := map[string]bool{
		".pdf": true,
		".txt": true,
	}

	// Destination directory for saving files
	destinationDir := "./uploads/"

	// Create the destination directory if it doesn't exist
	if _, err := os.Stat(destinationDir); os.IsNotExist(err) {
		os.Mkdir(destinationDir, os.ModePerm)
	}

	// Iterate over the uploaded files
	var fileName []dto.UploadResponse
	for _, headers := range r.MultipartForm.File {
		for _, header := range headers {
			// Check if the file has an allowed extension
			ext := strings.ToLower(filepath.Ext(header.Filename))
			if !allowedExtensions[ext] {
				err = errors.New("File extention not allowed")
				return nil, err
			}

			// Open the source file
			file, err := header.Open()
			if err != nil {
				log.Error(err)
				return nil, err
			}
			defer file.Close()

			// Create the destination file
			destinationFile, err := os.Create(destinationDir + header.Filename)
			if err != nil {
				log.Error(err)
				return nil, err
			}
			defer destinationFile.Close()

			// Copy the contents from the source file to the destination file
			_, err = io.Copy(destinationFile, file)
			if err != nil {
				log.Error(err)
				return nil, err
			}

			// fmt.Fprintf(w, "File %s uploaded successfully\n", header.Filename)
			fileName = append(fileName, dto.UploadResponse{
				FileName: header.Filename,
			})
		}
	}

	return fileName, err
}
