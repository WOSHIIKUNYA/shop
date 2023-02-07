package modle

type Good struct {
	Name    string `json:"name"`
	Id      int    `json:"int"`
	Kind    string `json:"kind"`
	Comment int    `json:"comment"`
	Picture string `json:"picture"`
	Price   int    `json:"Price"`
	Owner   string `json:"owner"`
}
type AddCart struct {
	Phone string `json:"Phone"`
	Id    int    `json:"id"`
}
type Cart struct {
	Id      int    `json:"id"`
	Phone   string `json:"Phone"`
	Price   int    `json:"price"`
	Picture string `json:"picture"`
	Name    string `json:"name"`
}
type DeleteCart struct {
	Id int `json:"id"`
}
type ClearCart struct {
	Id    int    `json:"id"`
	Phone string `json:"phone"`
}
