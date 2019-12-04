package app

type Error struct {
	Message      string        `json:"message"`
	Details      interface{}   `json:"details"`
	Description  string        `json:"description"`
	Code         string        `json:"code"`
}

