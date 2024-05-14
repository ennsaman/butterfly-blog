package model

import "time"

type Page struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Label      string    `json:"label"`
	Cover      string    `json:"cover"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
