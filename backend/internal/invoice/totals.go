package invoice

type Item struct {
	Quantity  int64
	RateCents int64
}

type Totals struct {
	SubtotalCents int64
	TaxCents      int64
	TotalCents    int64
}

func ComputeTotals(items []Item, taxRateBps int64) Totals {
	var subtotal int64
	for _, item := range items {
		subtotal += item.Quantity * item.RateCents
	}

	tax := subtotal * taxRateBps / 10000
	total := subtotal + tax

	return Totals{
		SubtotalCents: subtotal,
		TaxCents:      tax,
		TotalCents:    total,
	}
}
