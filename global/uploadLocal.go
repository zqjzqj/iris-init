package global

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"io"
	"net/http"
	"os"
	"strings"
)

func UploadLocalByUrl(url, path, filename string) (filePath string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()
	if filename == "" {
		filename = Md5(uuid.New().String()) + GetFileSuffix(url)
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
		filename = Md5(uuid.New().String()) + GetFileSuffix(info.Filename)
	}
	filename = strings.TrimRight(path, "/") + "/" + filename
	err = UploadLocalByReader(file, filename)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func UploadLocalByReader(file io.Reader, filename string) error {
	saveFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer func() { _ = saveFile.Close() }()
	_, err = io.Copy(saveFile, file)
	if err != nil {
		return err
	}
	return nil
}
