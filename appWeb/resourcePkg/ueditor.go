package resourcePkg

type Ueditor struct {
}

func (Ueditor) GetCss() []string {
	return nil
}

func (Ueditor) GetJs() []string {
	return []string{
		"/static/ueditor/ueditor.config.js",
		"/static/ueditor/ueditor.all.js",
		"/static/ueditor/ueditor.init.js",
	}
}
