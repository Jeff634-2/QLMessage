package engine

/*
	请求的内容
	解析urld的函数
*/

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []QLItem
}

type QLItem struct {
	QuanLiangName           string
	QLImage                 string
	QuanLiangPublishData    string
	QuanLiangPublishLearner string
}
