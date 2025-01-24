package db

import (
	"budgetsyncgo/models"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
	"log"
	"os"
	"time"
)

const (
	fetchTransactionsQuery = "SELECT t.transaction_pk, t.name, t.amount, c.name AS category_name, t.date_created FROM transactions t JOIN categories c ON t.category_fk = c.category_pk WHERE t.date_created >= ?"
)

type Sqlite3Handler struct {
	db *sql.DB
}

// NewSqlite3Handler initializes the database connection
func NewSqlite3Handler(dbFile string) (*Sqlite3Handler, error) {

	if _, err := os.Stat(dbFile); dbFile != ":memory:" && err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("database file %s does not exist: %w", dbFile, err)
		}
		return nil, fmt.Errorf("error checking database file: %w", err)
	}

	db, _ := sql.Open("sqlite3", dbFile)

	return &Sqlite3Handler{db: db}, nil
}

// FetchTransactionsStr retrieves transactions by converting a date string to time.Time
func (h *Sqlite3Handler) FetchTransactionsStr(dateFilter string) ([]models.Transaction, error) {
	parsedDate, err := parseDate(dateFilter)
	if err != nil {
		return nil, fmt.Errorf("failed to parse dateFilter (%s): %w", dateFilter, err)
	}
	return h.FetchTransactions(parsedDate)
}

// parseDate is a reusable helper function for parsing a date string into time.Time
func parseDate(dateStr string) (int64, error) {
	parsedTime, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return 0, fmt.Errorf("failed to parse date (%s): %w", dateStr, err)
	}
	return parsedTime.Unix(), nil
}

// FetchTransactions retrieves transactions filtered by date
func (h *Sqlite3Handler) FetchTransactions(dateFilter int64) ([]models.Transaction, error) {
	rows, err := h.db.Query(fetchTransactionsQuery, dateFilter)
	if err != nil {
		return nil, fmt.Errorf("failed to query TRANSACTIONS table: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Panic(err)
		}
	}(rows)

	var transactions []models.Transaction
	for rows.Next() {
		var txn models.Transaction
		var dateCreated int64 // Store Unix timestamp from database

		if err := rows.Scan(&txn.Id, &txn.Description, &txn.Amount, &txn.Category, &dateCreated); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Parse date string into time.Time
		txn.Date = time.Unix(dateCreated, 0).UTC() // assuming date is formatted as YYYY-MM-DD

		transactions = append(transactions, txn)
	}
	return transactions, nil
}
