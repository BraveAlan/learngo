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
		name := string(m[2])
		//result.Items = append(result.Items, "User "+name) // City
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]), // Url
			ParserFunc: func(c []byte) engine.ParseResult {
				//return ParseProfile(c, string(m[2])) // 这里m的作用域是这个for语句，{}里面会共用这个m，这导致所有人的名字是同一个
				return ParseProfile(c, name)
			},
		})
	}
	return result
}
