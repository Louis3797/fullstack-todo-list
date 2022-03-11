package models

type Todo struct {
	ID        string `redis:"id" json:"id" binding:"required"`
	Text      string `redis:"text" json:"text" binding:"required"`
	Status    bool   `redis:"status" json:"status" binding:"required"`
	Until     int    `redis:"until" json:"until" binding:"required"`
	CreatedAt int    `redis:"created_at" json:"created_at" binding:"required"`
}
