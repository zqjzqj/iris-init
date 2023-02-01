package ueditor

import (
	"errors"
	"iris-init/ueditor/storage"
	"net/http"
)

const (
	SUCCESS       string = "SUCCESS" //上传成功标记，在UEditor中内不可改变，否则flash判断会出错
	ERROR         string = "ERROR"
	NO_MATCH_FILE string = "no match file"
)

type Service struct {
	rootPath string // 项目根目录
	uploader UploaderInterface
}

func NewService(uploaderObj UploaderInterface, listObj ListInterface, rootPath string, configFilePath string) (serv Service, err error) {
	if uploaderObj == nil {
		// 没有注入uploader接口，则使用默认的方法
		uploaderObj = &Uploader{
			Storage: &storage.LocalFile{},
		}
	}
	_ = uploaderObj.SetRootPath(rootPath)

	if listObj == nil {
		listObj = &List{}
	}
	_ = listObj.SetRootPath(rootPath)

	serv = Service{
		rootPath: rootPath,
		uploader: uploaderObj,
	}

	if configFilePath != "" {
		// 加载默认配置
		err = loadConfigFromFile(configFilePath)
	}

	return
}

/*
*
上传图片
*/
func (serv Service) Uploadimage(r *http.Request) (res ResFileInfoWithState, err error) {
	fieldName := GloabConfig.ImageFieldName
	res, err = serv.uploadFile(r, fieldName, UploaderParams{
		PathFormat: GloabConfig.ImagePathFormat,
		MaxSize:    GloabConfig.ImageMaxSize,
		AllowFiles: GloabConfig.ImageAllowFiles,
	})
	defer func() {
		_ = r.Body.Close()
	}()
	return
}

/*
*
上传涂鸦
*/
func (serv Service) UploadScrawl(r *http.Request) (res ResFileInfoWithState, err error) {
	params := UploaderParams{
		PathFormat: GloabConfig.ScrawlPathFormat,
		MaxSize:    GloabConfig.ScrawlMaxSize,
		AllowFiles: GloabConfig.ImageAllowFiles,
		OriName:    "scrawl.png",
	}

	fieldName := GloabConfig.ScrawlFieldName
	data := r.PostFormValue(fieldName)
	_ = serv.uploader.SetParams(params)

	res = ResFileInfoWithState{}
	fileInfo, err := serv.uploader.UpBase64(params.OriName, data)
	defer func() {
		_ = r.Body.Close()
	}()
	if err == nil {
		res.fromResFileInfo(fileInfo)
		res.State = SUCCESS
		return
	}
	res.State = err.Error()
	return
}

/*
*
上传视频
*/
func (serv Service) UploadVideo(r *http.Request) (res ResFileInfoWithState, err error) {
	fieldName := GloabConfig.VideoFieldName
	res, err = serv.uploadFile(r, fieldName, UploaderParams{
		PathFormat: GloabConfig.VideoPathFormat,
		MaxSize:    GloabConfig.VideoMaxSize,
		AllowFiles: GloabConfig.VideoAllowFiles,
	})
	defer func() {
		_ = r.Body.Close()
	}()
	return
}

/*
*
上传文件
*/
func (serv Service) UploadFile(r *http.Request) (res ResFileInfoWithState, err error) {
	fieldName := GloabConfig.FileFieldName
	res, err = serv.uploadFile(r, fieldName, UploaderParams{
		PathFormat: GloabConfig.FilePathFormat,
		MaxSize:    GloabConfig.FileMaxSize,
		AllowFiles: GloabConfig.FileAllowFiles,
	})
	defer func() {
		_ = r.Body.Close()
	}()
	return
}

func (serv Service) uploadFile(r *http.Request, fieldName string, params UploaderParams) (fileInfo ResFileInfoWithState, err error) {
	file, fileHeader, err := r.FormFile(fieldName)
	if err != nil {
		return
	}

	fileInfo = ResFileInfoWithState{}

	_ = serv.uploader.SetParams(params)

	resFileInfo, err := serv.uploader.UpFile(file, fileHeader)
	if err == nil {
		fileInfo.fromResFileInfo(resFileInfo)
		fileInfo.State = SUCCESS
		return
	}
	fileInfo.State = err.Error()
	return
}

/*
*
读取配置信息
*/
func (serv Service) Config() (cnf *Config) {
	return GloabConfig
}

/*
*
获取图片列表
*/
func (serv Service) ListImage(listFileItem *ListFileItem, start int, size int) {
	list := &List{
		RootPath: serv.rootPath,
		Params: ListParams{
			AllowFiles: GloabConfig.ImageManagerAllowFiles,
			ListSize:   GloabConfig.ImageManagerListSize,
			Path:       GloabConfig.ImageManagerListPath,
		},
	}

	fileList, _ := list.GetFileList(start, size)
	if len(fileList) > 0 {
		listFileItem.State = SUCCESS
		listFileItem.List = fileList
		listFileItem.Total = len(fileList)
		listFileItem.Start = start
	} else {
		listFileItem.State = NO_MATCH_FILE
		listFileItem.List = fileList
		listFileItem.Total = 0
		listFileItem.Start = start
	}
}

/*
*
获取文件列表
*/
func (serv Service) Listfile(listFileItem *ListFileItem, start int, size int) {
	list := &List{
		RootPath: serv.rootPath,
		Params: ListParams{
			AllowFiles: GloabConfig.FileManagerAllowFiles,
			ListSize:   GloabConfig.FileManagerListSize,
			Path:       GloabConfig.FileManagerListPath,
		},
	}

	fileList, _ := list.GetFileList(start, size)
	if len(fileList) > 0 {
		listFileItem.State = SUCCESS
		listFileItem.List = fileList
		listFileItem.Total = len(fileList)
		listFileItem.Start = start
	} else {
		listFileItem.State = NO_MATCH_FILE
		listFileItem.List = fileList
		listFileItem.Total = 0
		listFileItem.Start = start
	}
}

/*
*
从远程拉取图片
*/
func (serv Service) CatchImage(r *http.Request) (listRes ResFilesInfoWithStates, err error) {
	_ = serv.uploader.SetParams(UploaderParams{
		PathFormat: GloabConfig.CatcherPathFormat,
		MaxSize:    GloabConfig.CatcherMaxSize,
		AllowFiles: GloabConfig.CatcherAllowFiles,
		OriName:    "remote.png",
	})

	fieldName := GloabConfig.CatcherFieldName

	err = r.ParseForm()
	if err != nil {
		err = errors.New("form parse error")
		return
	}

	source, _ := r.PostForm[fieldName+"[]"]

	filesInfos := make([]ResFileInfo, 0)
	for _, imgurl := range source {
		fileInfo, _ := serv.uploader.SaveRemote(imgurl)
		filesInfos = append(filesInfos, fileInfo)
	}

	if len(filesInfos) > 0 {
		listRes.State = SUCCESS
		listRes.List = filesInfos
		return
	}
	listRes.State = ERROR
	return
}
