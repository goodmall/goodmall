package demo

type Todo struct {
	Id          int    `json:"id"` // int32
	Created     int    `json:"created"`
	Status      string `json:"status"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
