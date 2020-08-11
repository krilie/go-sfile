package file_util

import (
	"bytes"
	"context"
	"io"
	"net/http"
)

// GetContentType2 从file中读出最多512字节用于确定类型 并所回一个新的reader供后续使用
func GetContentType(ctx context.Context, file io.Reader) (mType string, newReader io.Reader, err error) {
	decByte := make([]byte, 512)
	if readLen, err := file.Read(decByte); err != nil {
		return "", file, err
	} else {
		contentType := http.DetectContentType(decByte[0:readLen])
		return contentType, io.MultiReader(bytes.NewReader(decByte[0:readLen]), file), nil
	}
}
