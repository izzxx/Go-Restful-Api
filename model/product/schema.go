package product

import "time"

type Product struct {
	Id         string
	Name       string
	Price      float64
	Quantity   uint16
	Created_At time.Time
}
