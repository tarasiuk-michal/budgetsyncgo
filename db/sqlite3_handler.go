package db

import (
	"budgetsyncgo/models"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
	"log"
	"os"
	"time"
)

const (
	fetchTransactionsQuery = "SELECT id, description, amount, category, date FROM transactions WHERE date >= ?"
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

// FetchTransactions retrieves transactions filtered by date
func (h *Sqlite3Handler) FetchTransactions(dateFilter time.Time) ([]models.Transaction, error) {
	if dateFilter.IsZero() {
		return nil, errors.New("dateFilter cannot be zero")
	}
	rows, err := h.db.Query(fetchTransactionsQuery, dateFilter.Format(time.DateOnly))
	if err != nil {
		return nil, err
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
		var dateStr string // Intermediate storage for the 'date' column

		if err := rows.Scan(&txn.Id, &txn.Description, &txn.Amount, &txn.Category, &dateStr); err != nil {
			return nil, err
		}

		// Parse date string into time.Time
		txn.Date, err = time.Parse(time.DateOnly, dateStr) // assuming date is formatted as YYYY-MM-DD
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, txn)
	}
	return transactions, nil
}
