package store

type discount struct {
	threshold, rate float64
}

func (d *discount) calculateDiscount(price float64) float64 {
	discount := 0.0
	if price > d.threshold {
		discount = price * d.rate
	}
	return discount
}
