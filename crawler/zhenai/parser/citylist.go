package parser

import (
	"regexp"

	"imooc.com/sb/learngo/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, string(m[2])) // City
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]), // Url
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
