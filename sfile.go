package s_file

import (
	"context"
	"github.com/krilie/s-file/file_util"
	"io"
	"os"
)

type SFile struct {
	Dir string // 保存路径
}

func NewSFile(Dir string) *SFile {
	err := os.MkdirAll(Dir, 0777)
	if err != nil {
		panic(err)
	}
	return &SFile{Dir: Dir}
}

func (a *SFile) SaveFile(ctx context.Context, name string, f io.Reader) (content, key string, err error) {
	return file_util.SaveFile(ctx, a.Dir, name, f)
}
func (a *SFile) DeleteFile(ctx context.Context, key string) error {
	return file_util.RemoveFile(ctx, a.Dir, key)
}
