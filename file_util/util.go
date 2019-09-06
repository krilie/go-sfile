package file_util

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
)

func GetFileContent(ctx context.Context, file io.ReadSeeker) (content string, err error) {
	decByte := make([]byte, 512)
	if _, err := file.Read(decByte); err != nil {
		return "", err
	}
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", err
	}
	contentType := http.DetectContentType(decByte)
	return contentType, nil
}

func GetFileSha256(ctx context.Context, file io.ReadSeeker) (sha string, err error) {
	defer func() {
		_, err = file.Seek(0, io.SeekStart)
	}()
	hash := sha256.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	sum := hash.Sum(nil)
	return hex.EncodeToString(sum), nil
}
