package entity

import "time"

type Transaction struct {
	Amount    float32
	Category  string
	CreatedAt time.Time
}
