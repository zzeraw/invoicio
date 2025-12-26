package invoice

import (
	"context"
	"database/sql"
)

func NextInvoiceNumber(ctx context.Context, db *sql.DB) (int64, error) {
	var number int64
	err := db.QueryRowContext(ctx, "SELECT nextval('invoice_number_seq')").Scan(&number)
	if err != nil {
		return 0, err
	}
	return number, nil
}
