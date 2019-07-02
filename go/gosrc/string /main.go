package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func getLocation() {
	location := "周浦灾备机房-3楼302机房-L02-09~10"
	usite := "09~10"

	ind := strings.Index(location, usite)
	if ind > 1 {
		fmt.Printf("%s\n", location[0:ind-1])
	}

}

func other() {
	loc := "::::周浦灾备机房3楼302机房L02"
	locn := "周浦灾备机房-3楼302机房-L02"
	locnm := strings.Replace(locn, "-", "", -1)
	index := strings.IndexAny(loc, locnm)
	usite := loc[index+len(locnm):]

	fmt.Printf("locnm: %s, index: %d, usite: %s", locnm, index, usite)

	list := strings.Split("周浦灾备机房-3楼302机房-L02-09~10", "-")
	lenl := len(list)
	if lenl == 4 { // 4级结构说明包含U位信息，删除U位，进行位置校验
		list = list[:len(list)-1] // 去掉对U位校验
	}
	fmt.Printf("%v\n", list)

	getLocation()
}

var (
	locations = []string{"马场电信机房-3楼-E10-23-IS5347-0140",
		"马场电信机房-3楼-E10-05~35",
		"马场电信机房-3楼-E10-05~11-1#",
		"马场电信机房-3楼-E10-IS5347-0",
		"3楼",
		"C09",
		"马场电信机房-3楼-E10-2-~22",
		"马场电信机房-3楼-E10-08~09-IS5347-0216"}
	idc     = "马场电信机房"
	room    = []string{"3楼", "3楼1号", "3楼I区"}
	cabinet = []string{"NL01", "I3", "C12", "3fk11b", "E10", "NL01"}
)

func transString(src string) string {
	result, _ := simplifiedchinese.GB18030.NewDecoder().Bytes([]byte(src))
	return string(result)
}

func getUsite(location, idc, room, cabinet string) (usite string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("location: %s, idc: %s, room: %s, cabinet: %s\n", location, idc, room, cabinet)
		}
	}()
	// fmt.Printf("===============\n")
	// fmt.Printf("all:%s\n", location)
	// 提取数据中心
	if strings.Contains(location, idc) {
		indexIdc := strings.IndexAny(location, idc)
		if indexIdc >= 0 {
			location = location[len(idc):]

			if len(location) > 1 {
				if location[0] == '-' {
					location = location[1:]
				}
			}
		}
	}
	// fmt.Printf("room:%s\n", location)
	// 提取机房
	if strings.Contains(location, room) {
		indexRoom := strings.IndexAny(location, room)
		if indexRoom >= 0 {
			location = location[len(room):]
			if len(location) > 1 {
				if location[0] == '-' {
					location = location[1:]
				}
			}
		}
	}
	// fmt.Printf("cabinet: %s\n", location)
	// 提取机架
	if strings.Contains(location, cabinet) {
		indexCabinet := strings.IndexAny(location, cabinet)
		if indexCabinet >= 0 {
			location = location[len(cabinet):]
			if len(location) > 1 {
				if location[0] == '-' {
					location = location[1:]
				}
			}
		}
	}
	// fmt.Printf("usite: %s\n", location)

	if len(location) > 1 {
		if location[0] == '-' {
			location = location[1:]
		}
	}
	return location
}

func main() {
	var usite string
	for k := range locations {
		usite = getUsite(locations[k], "马场电信机房", "3楼", "C09")
		fmt.Printf("usite: %s\n", usite)
	}

	//fmt.Printf("%d\n", strings.IndexAny("马场电信机房-3楼-E10-23-IS5347-0140", "Josing"))
	//test := "马场电信机房"
	//fmt.Printf("%v\n", strings.Contains("马场电信机房C09", "10楼"))
}
