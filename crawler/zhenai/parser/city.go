package parser

import (
	"regexp"

	"imooc.com/sb/learngo/crawler/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9a-zA-z]+)" target="_blank">([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		//fmt.Printf("ParseCity %s", string(m[2]))
		result.Items = append(result.Items, "User "+string(m[2])) // City
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]), // Url
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, string(m[2]))
			},
		})
	}
	return result
}
