package models

import "time"

type List struct {
	Id         int64     `json:"id" db:"list_id"`
	Content    string    `json:"content" db:"content" `
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}
