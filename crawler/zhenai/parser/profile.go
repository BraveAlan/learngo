package parser

import (
	"log"
	"regexp"
	"strings"

	"imooc.com/sb/learngo/crawler/model"

	"imooc.com/sb/learngo/crawler/engine"
)

var ageRe = regexp.MustCompile(`<div.+?class="m-btn purple">([\d]+)岁</div>`)

var infoRe = regexp.MustCompile(`<div class="des f-cl".+?>(.+?)</div>`)
var nameRe = regexp.MustCompile(`<h1.+?class="nickName".+?>(.+?)</h1>`)

func ParseProfile(content []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	basicInfo := extractString(content, infoRe)
	basicInfoLabels := strings.Split(basicInfo, "|")
	if len(basicInfoLabels) != 6 {
		log.Printf("wrong basic info counts: %d", len(basicInfoLabels))
	} else {
		TrimSlice(basicInfoLabels)
		profile.House = basicInfoLabels[0]
		profile.Age = basicInfoLabels[1]
		profile.Education = basicInfoLabels[2]
		profile.Marriage = basicInfoLabels[3]
		profile.Height = basicInfoLabels[4]
		profile.Income = basicInfoLabels[5]
	}

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content) // 只获取第一个找到的串，match[0]是整个串，match[1]代表找到的东西
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

func TrimSlice(arr []string) {
	for i, str := range arr {
		arr[i] = strings.Trim(str, " ")
	}
}
