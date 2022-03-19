package models

type Todo struct {
	ID        string `redis:"id" json:"id"`
	Text      string `redis:"text" json:"text"`
	Status    bool   `redis:"status" json:"status"`
	CreatedAt int    `redis:"created_at" json:"created_at"`
}
