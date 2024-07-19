package global

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	idvalidator "github.com/guanguans/id-validator"
	"iris-init/sErr"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

var local *time.Location

type DuplicateType interface {
	NumberIntType | float32 | float64 | string
}

type NumberIntType interface {
	int | uint | int64 | uint64 | int8 | uint8 | int32 | uint32 | int16 | uint16
}

type ShuffleType interface {
	map[string]interface{} | NumberIntType | string
}

// 上传文件序号
var UploadFileNum uint64

func IsZeroValue(i interface{}) bool {
	return reflect.DeepEqual(i, reflect.Zero(reflect.TypeOf(i)).Interface())
}

func GetHttpBodyBytes(_url string) ([]byte, error) {
	resp, err := http.Get(_url)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	return io.ReadAll(resp.Body)
}

func RemoveFile(path string, level int) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	_path := path
	for i := 0; i < level; i++ {
		dir := filepath.Dir(_path)
		if !IsDirEmpty(dir) {
			break
		}
		err = os.RemoveAll(dir)
		if err != nil {
			return err
		}
		_path = dir
	}
	return nil
}

func IsDirEmpty(dir string) bool {
	entries, _ := os.ReadDir(dir)
	return len(entries) == 0
}

func SizeToString(size int64) string {
	const (
		B  = 1
		KB = 1024 * B
		MB = 1024 * KB
		GB = 1024 * MB
	)

	switch {
	case size < KB:
		return fmt.Sprintf("%d B", size)
	case size < MB:
		return fmt.Sprintf("%.2f KB", float64(size)/float64(KB))
	case size < GB:
		return fmt.Sprintf("%.2f MB", float64(size)/float64(MB))
	default:
		return fmt.Sprintf("%.2f GB", float64(size)/float64(GB))
	}
}

func IsNumber(v any) bool {
	var ref reflect.Type
	_v, ok := v.(reflect.Type)
	if ok {
		ref = _v
	} else {
		ref = reflect.TypeOf(v)
	}
	kind := ref.Kind()
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func GetNewFilename(oldFilename string) string {
	i := atomic.AddUint64(&UploadFileNum, 1)
	return fmt.Sprintf("%s%d_%s", time.Now().Format(DateTimeFormatStrCompact), i, oldFilename)
}

func GetFileSuffix(filename string) string {
	_suffix := strings.Split(filename, ".")
	_suffixLen := len(_suffix)
	if _suffixLen > 1 {
		return "." + _suffix[_suffixLen-1]
	}
	return ""
}

// StringFirstLower 字符串首字母小写
func StringFirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
func StringFirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
func GeneratePassword(pwd string) (password, salt string) {
	salt = RandStringRunes(10)
	return PwdPlaintext2CipherText(pwd, salt), salt
}

func CheckPassword(pwd string) error {
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			continue
		} else {
			return sErr.New("密码至少要包含数字大写小写字母")
		}
	}
	if StrLen(pwd) < 8 {
		return sErr.New("密码长度不能小于8")
	}
	return nil
}

// 驼峰转蛇形
func SnakeString(s string) string {
	if s == "ID" || s == "Id" {
		return "id"
	}
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		// or通过ASCII码进行大小写的转化
		// 65-90（A-Z），97-122（a-z）
		//判断如果字母为大写的A-Z就在前面拼接一个_
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	//ToLower把大写字母统一转小写 并且替换错误的ID转化
	return strings.ReplaceAll(strings.ToLower(string(data[:])), "i_d", "id")
}

func Day2YearMonthStrCh(_day uint, round bool) string {
	if _day == 0 {
		return ""
	}
	year := _day / 365
	month := (_day % 365) / 30
	day := (_day % 365) % 30
	if year > 0 {
		if (month == 0 && day == 0) || round {
			if year == 1 {
				return "年"
			}
			return fmt.Sprintf("%d年", year)
		}
		if month == 0 {
			return fmt.Sprintf("%d年零%d天", year, day)
		}
		if day == 0 {
			return fmt.Sprintf("%d年零%d个月", year, month)
		}
		return fmt.Sprintf("%d年零%d个月零%d天", year, month, day)
	}
	if month > 0 {
		if month <= 3 {
			return fmt.Sprintf("%d天", month*30+day)
		}
		if day == 0 || round {
			if month == 1 {
				return "月"
			}
			return fmt.Sprintf("%d个月", month)
		}
		return fmt.Sprintf("%d个月零%d天", month, day)
	}
	if day == 1 {
		return "天"
	}
	return fmt.Sprintf("%d天", day)
}

func Sce2TimeCh(sce uint, showSec bool) string {
	h := sce / 3600
	i := (sce % 3600) / 60
	s := (sce % 3600) % 60
	r := ""
	if h > 0 {
		r = fmt.Sprintf("%d小时", h)
	}
	if i > 0 {
		r += fmt.Sprintf("%d分钟", i)
	}
	if showSec && s > 0 {
		r += fmt.Sprintf("%d秒", s)
	}
	return r
}

func Sce2TimeStr(sce uint, ellipsisHour bool) string {
	h := sce / 3600
	i := (sce % 3600) / 60
	s := (sce % 3600) % 60
	var hh, ii, ss string
	if ellipsisHour && h > 0 {
		if h < 10 {
			hh = fmt.Sprintf("0%d", h)
		} else {
			hh = fmt.Sprintf("%d", h)
		}
	}
	if i < 10 {
		ii = fmt.Sprintf("0%d", i)
	} else {
		ii = fmt.Sprintf("%d", i)
	}
	if s < 10 {
		ss = fmt.Sprintf("0%d", s)
	} else {
		ss = fmt.Sprintf("%d", s)
	}
	if ellipsisHour && h == 0 {
		return fmt.Sprintf("%s:%s", ii, ss)
	}
	return fmt.Sprintf("%s:%s:%s", hh, ii, ss)
}

func Shuffle[T ShuffleType](slice []T) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

// 对长度不足n的数字前面补0
func SupNumber[T DuplicateType](i T, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}

func Bytes2UUID(b []byte) (uuid.UUID, error) {
	_uuid := uuid.New()
	if err := _uuid.UnmarshalBinary(b); err != nil {
		return _uuid, err
	}
	return _uuid, nil
}

func UUIDStr2Bytes(_uuidStr string) ([]byte, error) {
	_uuid := uuid.New()
	err := _uuid.Scan(_uuid)
	if err != nil {
		return nil, err
	}
	return _uuid.MarshalBinary()
}

func GenerateToken(complex uint) string {
	return Md5(fmt.Sprintf("%s%s", uuid.New().String(), RandStringRunes(int(complex))))
}

func Decimal(value float64, prec int) float64 {
	value, _ = strconv.ParseFloat(strconv.FormatFloat(value, 'f', prec, 64), 64)
	return value
}

func Camel2String(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func InSlice[T DuplicateType](value T, element []T) bool {
	for _, item := range element {
		if item == value {
			return true
		}
	}
	return false
}

func CheckPhone(phone string) bool {
	regular := "^((19[0-9])|(13[0-9])|(14[0-9])|(15[0-9])|(17[0-9])|(18[0-9])|(16[0-9])|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
	/*	pl := StrLen(phone)
		if phone == "" || pl > 15 || pl < 11 {
			return false
		}
		return true*/
}

func CheckIDCard(IDCard string) bool {
	return idvalidator.IsValid(IDCard, false)
}

// 移除重复元素
func RemoveDuplicateElement[T DuplicateType](element []T) []T {
	result := make([]T, 0, len(element))
	temp := map[T]struct{}{}
	for _, item := range element {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// 移除指定元素
func RemoveElement[T DuplicateType](item T, items []T) ([]T, int) {
	num := 0
	for i := 0; i < len(items); i++ {
		if items[i] == item {
			num++
			items = items[:i+copy(items[i:], items[i+1:])]
		}
	}
	return items, num
}

func StrArrToUintArr(element []string) []uint64 {
	elementUint64 := make([]uint64, 0, len(element))
	for _, v := range element {
		i, err := strconv.ParseUint(v, 10, 64)
		if err == nil {
			elementUint64 = append(elementUint64, i)
		}
	}
	return elementUint64
}

func Int8Arr2IntArr(r []uint8) []int {
	rr := make([]int, 0, len(r))
	for _, v := range r {
		rr = append(rr, int(v))
	}
	return rr
}

func NumberArrToStrArr[T NumberIntType](element []T) []string {
	elementString := make([]string, 0, len(element))
	for _, v := range element {
		elementString = append(elementString, fmt.Sprintf("%d", v))
	}
	return elementString
}

func StrLen(str string) int {
	return strings.Count(str, "") - 1
}

func PwdPlaintext2CipherText(pwd string, salt string) string {
	pwd = salt + "{_}" + pwd + "{_}" + salt
	has := md5.Sum([]byte(pwd))
	return fmt.Sprintf("%x", has)
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateRangeNum(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Int63n(max-min) + min
	return randNum
}

func Hour2Unix(hour string) (time.Time, error) {
	return time.ParseInLocation(time.DateTime, time.Now().Format(time.DateOnly)+" "+hour, GetLocalTime())
}

func Md5(s string) string {
	data := []byte(s)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func HmacSHA1(secretToken, payloadBody string) []byte {
	mac := hmac.New(sha1.New, []byte(secretToken))
	sha1.New()
	mac.Write([]byte(payloadBody))
	return mac.Sum(nil)
}

func Json2Map(j string) map[string]interface{} {
	r := make(map[string]interface{})
	_ = json.Unmarshal([]byte(j), &r)
	return r
}

func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func RandFloats(min, max float64, n int) float64 {
	rand.Seed(time.Now().UnixNano())
	res := min + rand.Float64()*(max-min)
	res, _ = strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(n)+"f", res), 64)
	return res
}

func SubString(s string, length int) string {
	ss := []rune(s)
	if len(ss) <= length {
		return s
	}
	return string(ss[:length])
}

func EllipsisString(s string, length int) string {
	ss := []rune(s)
	if len(ss) <= length {
		return s
	}
	return string(ss[:length]) + "..."
}

func Str2Time(layout, val string) (time.Time, error) {
	return time.ParseInLocation(layout, val, GetLocalTime())
}
func GetLastDayOfLastMonth(_t time.Time) time.Time {
	return _t.AddDate(0, 1, -_t.Day())
}

func GetFirstDateOfYear() time.Time {
	_now := time.Now()
	return time.Date(_now.Year(), 1, 1, 0, 0, 0, 0, GetLocalTime())
}

func GetLocalTime() *time.Location {
	if local == nil {
		local, _ = time.LoadLocation("Local")
	}
	return local
}

func GetFirstDateOfWeek() time.Time {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, GetLocalTime()).
		AddDate(0, 0, offset)
}

func HideStar(str string) (result string) {
	if str == "" {
		return "***"
	}
	if strings.Contains(str, "@") {
		res := strings.Split(str, "@")
		if len(res[0]) < 3 {
			resString := "***"
			result = resString + "@" + res[1]
		} else {
			res2 := Substr2(str, 0, 3)
			resString := res2 + "***"
			result = resString + "@" + res[1]
		}
		return result
	} else {
		reg := `^1[0-9]\d{9}$`
		rgx := regexp.MustCompile(reg)
		mobileMatch := rgx.MatchString(str)
		if mobileMatch {
			result = Substr2(str, 0, 3) + "****" + Substr2(str, 7, 11)
		} else {
			nameRune := []rune(str)
			lens := len(nameRune)

			if lens <= 1 {
				result = "***"
			} else if lens == 2 {
				result = string(nameRune[:1]) + "*"
			} else if lens == 3 {
				result = string(nameRune[:1]) + "*" + string(nameRune[2:3])
			} else if lens == 4 {
				result = string(nameRune[:1]) + "**" + string(nameRune[lens-1:lens])
			} else if lens > 4 {
				result = string(nameRune[:2]) + "***" + string(nameRune[lens-2:lens])
			}
		}
		return
	}
}

func Substr2(str string, start int, end int) string {
	rs := []rune(str)
	return string(rs[start:end])
}
