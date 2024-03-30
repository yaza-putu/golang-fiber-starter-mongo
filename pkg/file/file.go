package file

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
)

// DetectContentType to make sure the mimes of data
func DetectContentType(file multipart.File, allowMimes []string) bool {
	// get type MIME of file
	buffer := make([]byte, 512) // read 512 bytes to make sure MIME
	_, err := file.Read(buffer)
	if err != nil {
		return false
	}
	fileType := http.DetectContentType(buffer)
	// compare mime type
	for _, validType := range allowMimes {
		fmt.Println(fileType, validType)
		if strings.HasPrefix(fileType, validType) {
			return true
		}
	}
	return false
}
