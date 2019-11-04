package parser

import (
	"fmt"
	"gorobot/distributed/config"
	"gorobot/engine"
	"regexp"
)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/([0-9]+))"[^>]*>([^<]+)</a>`)
var cityURLRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

//ParserCity function
func ParseCity(
	contents []byte, _ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		fmt.Printf("got city:%s:%s:%s\n", m[1], m[2], m[3])
		//got city:http://album.zhenai.com/u/1604417986:1604417986:有缘就有希望
		result.Requests = append(result.Requests, engine.Request{
			GUrl:   string(m[1]),
			Parser: NewProfileParser(string(m[2]), string(m[3])),
		})
	}
	matches = cityURLRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			GUrl:   string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}

	return result
}
