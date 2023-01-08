package modle

type User struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
	Question string `json:"Question"`
	Answer   string `json:"Answer"`
	Phone    string `json:"Phone"`
}
type LoginUser struct {
	Phone    string `json:"Phone"`
	Password string `json:"Password"`
}
