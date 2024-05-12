package model

type User struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
