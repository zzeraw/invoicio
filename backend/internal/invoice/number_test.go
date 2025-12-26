package invoice

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestNextInvoiceNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"nextval"}).AddRow(int64(1001))
	mock.ExpectQuery("SELECT nextval\\('invoice_number_seq'\\)").WillReturnRows(rows)

	got, err := NextInvoiceNumber(context.Background(), db)
	if err != nil {
		t.Fatalf("NextInvoiceNumber: %v", err)
	}
	if got != 1001 {
		t.Fatalf("number = %d, want %d", got, 1001)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("expectations: %v", err)
	}
}
