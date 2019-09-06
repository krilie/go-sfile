package file_util

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

func SaveFile(ctx context.Context, dir, name string, file io.ReadSeeker) (content string, err error) {
	decByte := make([]byte, 512)
	if _, err := file.Read(decByte); err != nil {
		return "", err
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", err
	}
	contentType := http.DetectContentType(decByte)
	fullPath := fmt.Sprintf("%v%v%v", dir, os.PathSeparator, name)
	f, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	_, err = io.Copy(f, file)
	if err != nil {
		err2 := os.Remove(fullPath)
		return "", fmt.Errorf("%v %v", err, err2)
	}
	return contentType, nil
}

func RemoveFile(ctx context.Context, dir, name string) error {
	return os.Remove(fmt.Sprintf("%v%v%v", dir, os.PathSeparator, name))
}
