package engine

import (
	"gorobot/fetcher"
	//"log"
)

// Worker r Request
func Worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s", r.GUrl)

	body, err := fetcher.Fetch(r.GUrl)
	if err != nil {
		//log.Printf("Fetcher:error "+"fetching url %s:%v", r.GUrl, err)
		//2019/10/23 22:48:15 Fetching http://www.zhenai.com/zhenghun/dazu/xinxi
		return ParseResult{}, err
	}

	return r.ParserFunc(body, r.GUrl), nil
}
