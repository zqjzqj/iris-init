package global

// 这个包放置公共常量
const (
	IsYes = 1
	IsNo  = 0

	DateTimeFormatStr        = "2006-01-02 15:04:05"
	DateHisFormatStr         = "15:04:05"
	DateFormatStr            = "2006-01-02"
	DateYmFormatStr          = "2006-01"
	DateTzFormatStr          = "2006-01-02T15:04:05+08:00"
	DateTimeFormatStrCompact = "20060102150405"

	SexMan     = 1
	SexWoman   = 2
	SexUnknown = 0

	DefaultAvatar = "https://tcr-4w06clvb-1303139375.cos.ap-beijing.myqcloud.com/gzzkResources/2d9f1886d9b1f5de774d48baa519d132.png"
)

var SexDescMap = map[uint8]string{
	SexMan:     "男",
	SexWoman:   "女",
	SexUnknown: "未知",
}
