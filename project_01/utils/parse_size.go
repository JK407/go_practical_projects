package utils

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

// ParseSize
// @Description 解析传入 size 字符串，返回对应的字节数和单位
// @Author Oberl-Fitzgerald 2024-07-07 13:24:52
// @Param  size string
// @Return int64
// @Return string
func ParseSize(size string) (int64, string) {
	// 默认为 100 MB
	re, _ := regexp.Compile("[0-9]+")
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit)
	var byteNum int64 = 0
	switch unit {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * KB
	case "MB":
		byteNum = num * MB
	case "GB":
		byteNum = num * GB
	case "TB":
		byteNum = num * TB
	case "PB":
		byteNum = num * PB
	default:
		num = 0
		byteNum = 0
	}

	if num == 0 {
		log.Println("size is not valid,only support B,KB,MB,GB,TB,PB")
		num = 100
		byteNum = num * MB
		unit = "MB"
	}

	sizeStr := strconv.FormatInt(num, 10) + unit
	return byteNum, sizeStr
}
