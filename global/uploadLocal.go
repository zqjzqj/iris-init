package global

import (
	"bytes"
	"github.com/kataras/iris/v12"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UploadLocalByUrl(url, path, filename string) (filePath string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()
	if filename == "" {
		_url := strings.Split(strings.TrimRight(url, "/"), "/")
		filename = GetNewFilename(_url[len(_url)-1])
	}
	filename = strings.TrimRight(path, "/") + "/" + filename
	err = UploadLocalByReader(resp.Body, filename)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func UploadLocalByCtx(ctx iris.Context, field, path, filename string) (filePath string, err error) {
	file, info, err := ctx.FormFile(field)
	if err != nil {
		return "", err
	}
	defer func() { _ = file.Close() }()
	if filename == "" {
		//处理一下文件名 防止重复了
		filename = GetNewFilename(info.Filename)
		//filename = Md5(uuid.New().String()) + GetFileSuffix(info.Filename)
	}
	filename = strings.TrimRight(path, "/") + "/" + filename
	err = UploadLocalByReader(file, filename)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func UploadLocalByBytes(_byte []byte, filename string) error {
	return UploadLocalByReader(bytes.NewReader(_byte), filename)
}

func UploadLocalByReader(file io.Reader, filename string) error {
	fileDir := filepath.Dir(filename)
	if !FileExists(fileDir) {
		// 文件夹不存在，创建
		if err := os.MkdirAll(fileDir, 0766); err != nil {
			return err
		}
	}
	dstFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer func() { _ = dstFile.Close() }()
	_, err = io.Copy(dstFile, file)
	if err != nil {
		return err
	}
	return nil
}
