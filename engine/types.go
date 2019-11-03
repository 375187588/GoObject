package engine

type Item struct {
	URL     string
	Type    string
	ID      string
	Payload interface{}
}

type ParserFunc func(contents []byte, url string) ParseResult

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

//Request struct
type Request struct {
	GUrl   string
	Parser Parser
}

//{"ParseCityList",nil},{"ProfileParser",userName}

//ParseResult struct
type ParseResult struct {
	Requests []Request
	Items    []Item
}

type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(
	p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
