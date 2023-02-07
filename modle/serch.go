package modle

type SearchName struct {
	Name string `json:"name"`
}
type Goods struct {
	Name    string `json:"name"`
	Id      int    `json:"Id"`
	Kind    string `json:"kind"`
	Comment string `json:"comment"`
	Price   string `json:"Price"`
	Owner   string `json:"owner"`
	Picture string `json:"picture"`
}
type Tag struct {
	Kind string `json:"kind"`
}
type Id struct {
	Id int `json:"Id"`
}
