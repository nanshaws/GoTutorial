package cache

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

func ParseSize(size string) (int64, string) {
	//默认大小为100MB
	re, _ := regexp.Compile("[0-9]+")
	unit := string(re.ReplaceAll([]byte(size), []byte("")))
	num, _ := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	unit = strings.ToUpper(unit)
	var byteNum int64 = 0
	switch unit {
	case "B":
		byteNum = num
	case "KB":
		byteNum = KB * num
	case "MB":
		byteNum = MB * num
	case "GB":
		byteNum = GB * num
	case "TB":
		byteNum = TB * num
	case "PB":
		byteNum = PB * num
	default:
		num = 0

	}
	if num == 0 {
		log.Println("ParseSize 仅供支持B、KB、MB、GB、TB、PB")
		num = 100 * MB
		byteNum = num * MB
		unit = "MB"
	}
	sizeStr := strconv.FormatInt(num, 10) + unit
	return byteNum, sizeStr
}
