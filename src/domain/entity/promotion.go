package entity

import "time"

type Promotion struct {
	Id          int
	Isbn        string
	Discount    float64
	Enabled     bool
	CreatedDate time.Time
	UpdatedDate time.Time
}
