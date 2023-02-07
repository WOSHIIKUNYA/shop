package modle

type Comment struct {
	Name    string `json:"name"`
	Score   int    `json:"score"`
	GoodsId int    `json:"goods_id"`
	Body    string `json:"body"`
}
type GoodId struct {
	GoodsId int `json:"goods_id"`
}
type Writer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CommentAll struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Score   int    `json:"score"`
	GoodsId int    `json:"goods_id"`
	Body    string `json:"body"`
}
