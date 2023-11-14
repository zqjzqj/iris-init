package resourcePkg

type AreaSelect struct {
}

func (AreaSelect) GetCss() []string {
	return nil
}

func (AreaSelect) GetJs() []string {
	return []string{
		"areaSelect.js",
	}
}
