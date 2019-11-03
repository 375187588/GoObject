package engine

type ParserFunc func(
	contents []byte, url string) ParseResult

//Request struct
type Request struct {
	GUrl       string
	ParserFunc ParserFunc
}

//ParseResult struct
type ParseResult struct {
	Requests []Request
	Items    []Item
}

//NilParser byte
func NilParser([]byte) ParseResult {
	return ParseResult{}
}

type Item struct {
	URL     string
	Type    string
	ID      string
	Payload interface{}
}
