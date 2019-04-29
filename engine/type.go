package engine

type ParserFunc func(contents []byte, url string) ParseResult

type Request struct{
	Url        string
	ParserFunc ParserFunc
}

type ParseResult struct{
	Requests []Request
	Items []Item
}

type Item struct {
	Url string
	Type string //用来表示在elasticsearch里对应的Type
	Id string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}