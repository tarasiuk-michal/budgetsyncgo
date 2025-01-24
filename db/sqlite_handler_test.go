package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestFetchTransactions(t *testing.T) {
	// Initialize in-memory database
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Errorf("failed to open in-memory database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Panic(err)
		}
	}(db)

	// Create mock table and seed data
	setupTestDatabase(db, t)

	handler := &SqliteHandler{db: db}

	for _, tc := range FetchTransactionsTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Execute the FetchTransactions function
			got, err := handler.FetchTransactionsStr(tc.DateFilter)

			// Validate error outcome
			if (err != nil) != tc.WantErr {
				t.Fatalf("unexpected error: got %v, want error: %v", err, tc.WantErr)
			}

			// Validate the number of transactions
			assert.Equal(t, len(tc.WantDesc), len(got), "unexpected transaction count: got %d, want %d", len(got), len(tc.WantDesc))

			// Validate individual transaction descriptions
			for i, txn := range got {
				assert.Equal(t, tc.WantDesc[i], txn.Description, "unexpected transaction: got %s, want %s", txn.Description, tc.WantDesc[i])
			}
		})
	}
}

func TestNewSqliteHandler(t *testing.T) {
	for _, tc := range NewSqliteHandlerTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Call the NewSqliteHandler function
			handler, err := NewSqliteHandler(tc.DbFile)

			// Validate error outcome
			assert.Equal(t, tc.WantErr, err != nil, "unexpected error: got %v, want error: %v", err, tc.WantErr)

			if !tc.WantErr {
				// Validate handler properties
				assert.False(t, handler == nil, "handler is nil")
				assert.False(t, handler.db == nil, "handler initialized but database is nil")
			}
		})
	}
}

func setupTestDatabase(db *sql.DB, t *testing.T) {

	// Create mock categories table
	_, err := db.Exec(CreateCategoriesTableQuery)
	if err != nil {
		t.Fatalf("failed to create mock categories table: %v", err)
	}

	// Create mock transactions table
	_, err = db.Exec(CreateTransactionsTableQuery)
	if err != nil {
		t.Fatalf("failed to create mock transactions table: %v", err)
	}

	// Insert mock categories data
	_, err = db.Exec(InsertCategoriesDataQuery)
	if err != nil {
		t.Fatalf("failed to insert mock categories data: %v", err)
	}

	// Insert mock transactions data
	_, err = db.Exec(InsertTransactionsDataQuery)
	if err != nil {
		t.Fatalf("failed to insert mock transactions data: %v", err)
	}
}
