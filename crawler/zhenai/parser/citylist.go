package parser

import (
	"regexp"

	"imooc.com/sb/learngo/crawler/engine"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(contents []byte) engine.ParseResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 1
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2])) // City
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]), // Url
			ParserFunc: ParseCity,
		})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
