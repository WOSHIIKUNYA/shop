package modle

type User struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
	Question string `json:"Question"`
	Answer   string `json:"Answer"`
	Phone    string `json:"Phone"`
}
type Answer struct {
	Answer string `json:"Answer"`
}
type Login struct {
	Phone    string `json:"Phone"`
	Password string `json:"Password"`
}
type LoginUser struct {
	Phone    string `json:"Phone"`
	Name     string `json:"name"`
	Password string `json:"Password"`
	Picture  string `json:"picture"`
	Money    int    `json:"money"`
	Address  string `json:"address"`
}
type Phone struct {
	Phone string `json:"Phone"`
}
type Seek struct {
	Phone    string `json:"Phone"`
	Question string `json:"Question"`
	Answer   string `json:"Answer"`
}
