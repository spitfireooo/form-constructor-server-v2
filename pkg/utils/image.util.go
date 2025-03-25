package utils

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

func CheckContentType(contentTypeFile string, contentTypes ...string) error {
	if len(contentTypes) > 0 {
		for _, contentType := range contentTypes {
			if contentTypeFile == contentType {
				return nil
			}
		}
		return errors.New("not allowed type file")
	} else {
		return errors.New("not found content type to be checking")
	}
}

func GenerateFilename(filename string) string {
	return fmt.Sprintf("./static/uploads/%s_%s", uuid.New(), filename)
}
