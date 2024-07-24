package entity

import (
	"time"
)

type Todo struct {
	Id                 int32
	Title, Description string
	CreatedAt          time.Time
}
