package handlers

type Error struct {
	Message    string `json:"message"`
	StatusCode int `json:"code"`
	Status     string `json:"status"`
}
