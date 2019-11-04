package parser

import (
	"gorobot/distributed/config"
	"gorobot/engine"
	"regexp"
)

const cityListRe = `(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//ParseCityList cityList
func ParseCityList(
	contents []byte, _ string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1) //re.FindAllStringSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				GUrl: string(m[1]),
				Parser: engine.NewFuncParser(
					ParseCity, config.ParseCityList),
			})
	}

	return result
}
