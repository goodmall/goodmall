package demo

type Todo struct {
	Id          int32  `json:"id"`
	Created     int32  `json:"created"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
