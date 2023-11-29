package storage

import (
	"errors"
	"io"
	"iris-init/global"
	"iris-init/ueditor/ueditorCommon"
)

type LocalFile struct {
	BaseInterface
}

func (lf LocalFile) SaveFileFromLocalPath(srcPath string, dstAbsPath, dstRelativePath string) (url string, err error) {
	if !global.FileExists(srcPath) {
		err = errors.New(ueditorCommon.ERROR_FILE_STATE)
		return
	}
	return
}

/*
*
保存文件到本地
*/
func (lf LocalFile) SaveFile(srcFile io.Reader, srcFileSize int64, dstAbsPath, dstRelativePath string) (url string, err error) {
	_, err = global.UploadLocalByReader(srcFile, dstAbsPath)
	if err != nil {
		return
	}
	url = dstRelativePath
	return
}

/*
*
保存数据到本地
*/
func (lf LocalFile) SaveData(data []byte, dstAbsPath, dstRelativePath string) (url string, err error) {
	err = global.UploadLocalByBytes(data, dstAbsPath)
	if err != nil {
		return
	}
	url = dstRelativePath
	return
}
