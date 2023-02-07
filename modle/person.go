package modle

type Recharge struct {
	Phone string `json:"Phone"`
	Money int    `json:"money"`
}
type ChangeAddress struct {
	Phone   string `json:"Phone"`
	Address string `json:"address"`
}
type Change struct {
	Phone string `json:"Phone"`
	Name  string `json:"name"`
}
type See struct {
	Address string `json:"address"`
	Picture string `json:"picture"`
	Money   int    `json:"money"`
	Name    string `json:"name"`
}
