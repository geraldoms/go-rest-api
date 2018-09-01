package models

import (
	"time"
)

type ToDo struct {
	ID          int       `json:"id"`
	ToDo        string    `json:"todo"`
	Description string    `json:"description"`
	IsDone      bool      `json:"isdone"`
	CreateAt    time.Time `json:"createat"`
}
