package migrates

// 所需的迁移在写在这里 按照时间顺序排序
var MM = []MigrateInterface{
     {{- range $v := .AMigratesSlices}}
     {{$v}},
     {{- end}}
}
