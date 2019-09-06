package s_file

import (
	"context"
	"io"
	"os"
	"s-file/file_util"
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

func (a *SFile) SaveFile(ctx context.Context, name string, f io.ReadSeeker) (content, key string, err error) {
	return file_util.SaveFile(ctx, a.Dir, name, f)
}
func (a *SFile) DeleteFile(ctx context.Context, key string) error {
	return file_util.RemoveFile(ctx, a.Dir, key)
}
