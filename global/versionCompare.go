package global

import (
	"math"
	"strconv"
	"strings"
)

var (
	VersionBig   = 1
	VersionSmall = 2
	VersionEqual = 0
)

func Version2Number(v string, long int) int {
	verStrArr := spliteStrByNet(v)
	if len(verStrArr) != long {
		return 0
	}
	pow := math.Pow10(long)
	r := 0
	for _, val := range verStrArr {
		vNum, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return 0
		}
		r += int(pow) * int(vNum)
		pow = pow / 10
	}
	return r
}

//A与B比较 A大 return VersionBig A小 return VersionSmall 相等 return VersionEqual
func CompareStrVer(verA, verB string) int {
	verStrArrA := spliteStrByNet(verA)
	verStrArrB := spliteStrByNet(verB)
	lenStrA := len(verStrArrA)
	lenStrB := len(verStrArrB)

	//自动补齐长度
	if lenStrA != lenStrB {
		if lenStrA > lenStrB {
			for i := 0; i < lenStrA-lenStrB; i++ {
				verStrArrB = append(verStrArrB, "0")
			}
		} else {
			for i := 0; i < lenStrB-lenStrA; i++ {
				verStrArrA = append(verStrArrA, "0")
			}
		}
	}
	return _compareArrStrVer(verStrArrA, verStrArrB)
}

func _compareArrStrVer(verA, verB []string) int {
	for index, _ := range verA {
		littleResult := compareLittleVer(verA[index], verB[index])
		if littleResult != VersionEqual {
			return littleResult
		}
	}
	return VersionEqual
}

func compareLittleVer(verA, verB string) int {
	bytesA := []byte(verA)
	bytesB := []byte(verB)
	lenA := len(bytesA)
	lenB := len(bytesB)
	if lenA > lenB {
		return VersionBig
	}
	if lenA < lenB {
		return VersionSmall
	}
	return compareByBytes(bytesA, bytesB)
}

func compareByBytes(verA, verB []byte) int {
	for index, _ := range verA {
		if verA[index] > verB[index] {
			return VersionBig
		}
		if verA[index] < verB[index] {
			return VersionSmall
		}
	}
	return VersionEqual
}

func spliteStrByNet(strV string) []string {
	return strings.Split(strV, ".")
}
