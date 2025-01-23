package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"testing"
)

func TestFetchTransactions(t *testing.T) {
	// Initialize in-memory database
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Panic(err)
		}
	}(db)

	// Create mock table and seed data
	setupTestDatabase(db, t)

	handler := &Sqlite3Handler{db: db}

	for _, tc := range FetchTransactionsTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Execute the FetchTransactions function
			got, err := handler.FetchTransactions(tc.DateFilter)

			// Validate error outcome
			if (err != nil) != tc.WantErr {
				t.Fatalf("unexpected error: got %v, want error: %v", err, tc.WantErr)
			}

			// Validate the number of transactions
			if len(got) != len(tc.WantDesc) {
				t.Fatalf("unexpected transaction count: got %d, want %d", len(got), len(tc.WantDesc))
			}

			// Validate individual transaction descriptions
			for i, txn := range got {
				if txn.Description != tc.WantDesc[i] {
					t.Fatalf("unexpected transaction: got %s, want %s", txn.Description, tc.WantDesc[i])
				}
			}
		})
	}
}

func TestNewSqlite3Handler(t *testing.T) {
	for _, tc := range NewSqlite3HandlerTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			// Call the NewSqlite3Handler function
			handler, err := NewSqlite3Handler(tc.DbFile)

			// Validate error outcome
			if (err != nil) != tc.WantErr {
				t.Fatalf("unexpected error: got %v, want error: %v", err, tc.WantErr)
			}

			// Validate handler properties
			if handler != nil && handler.db == nil {
				t.Fatalf("handler initialized but database is nil")
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
