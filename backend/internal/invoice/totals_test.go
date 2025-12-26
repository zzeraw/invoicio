package invoice

import "testing"

func TestComputeTotals(t *testing.T) {
	items := []Item{
		{Quantity: 2, RateCents: 1000},
		{Quantity: 1, RateCents: 500},
	}

	got := ComputeTotals(items, 2000)
	if got.SubtotalCents != 2500 {
		t.Fatalf("subtotal = %d, want %d", got.SubtotalCents, 2500)
	}
	if got.TaxCents != 500 {
		t.Fatalf("tax = %d, want %d", got.TaxCents, 500)
	}
	if got.TotalCents != 3000 {
		t.Fatalf("total = %d, want %d", got.TotalCents, 3000)
	}
}
