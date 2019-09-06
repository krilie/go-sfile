package s_file

import (
	"context"
	"io"
	"s-file/file_util"
)

type SFile struct {
	Dir string // 保存路径
}

func NewSFile(Dir string) *SFile {
	return &SFile{Dir: Dir}
}

func (a *SFile) SaveFile(ctx context.Context, name string, f io.ReadSeeker) (content string, err error) {
	return file_util.SaveFile(ctx, a.Dir, name, f)
}
func (a *SFile) DeleteFile(ctx context.Context, name string) error {
	return file_util.RemoveFile(ctx, a.Dir, name)
}
