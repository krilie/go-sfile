package file_util

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
)

func SaveFile(ctx context.Context, dir, oriName string, file io.ReadSeeker) (content, key string, err error) {
	fileContent, err := GetFileContent(ctx, file)
	if err != nil {
		return "", "", err
	}
	key, err = GetFileSha256(ctx, file)
	if err != nil {
		return "", "", err
	}
	key = fmt.Sprintf("%s%s", key, url.PathEscape(oriName))
	fullPath := fmt.Sprintf("%s%s%s", dir, string(os.PathSeparator), key)
	f, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return "", "", err
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		err2 := os.Remove(fullPath)
		return "", "", fmt.Errorf("%v %v", err, err2)
	}
	return fileContent, key, nil
}

func RemoveFile(ctx context.Context, dir, key string) error {
	return os.Remove(fmt.Sprintf("%v%v%v", dir, os.PathSeparator, key))
}
