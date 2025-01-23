// Package db testdata.go
package db

import (
	"time"
)

// CreateTransactionsTableQuery Mock table creation query
const CreateTransactionsTableQuery = `
CREATE TABLE transactions (
	transaction_pk INTEGER PRIMARY KEY,
	name TEXT,
	amount REAL,
	category_fk INTEGER,
	date_created TEXT
)`

// CreateCategoriesTableQuery Mock table creation query
const CreateCategoriesTableQuery = `
CREATE TABLE categories (
	category_pk INTEGER PRIMARY KEY,
	name TEXT
)`

// InsertTransactionsDataQuery Mock table data insertion query
const InsertTransactionsDataQuery = `
INSERT INTO transactions (transaction_pk, name, amount, category_fk, date_created) VALUES
(1, 'Groceries', 50.0, 11, '2023-09-20'),
(2, 'Rent', 500.0, 22, '2023-10-01'),
(3, 'Salary', 2000.0, 33, '2023-10-05'),
(4, 'Thing', 300.0, 44, '2020-10-05 12:00:00')`

// InsertCategoriesDataQuery Mock table data insertion query
const InsertCategoriesDataQuery = `
INSERT INTO categories (category_pk, name) VALUES
(11, 'Food'),
(22, 'Housing'),
(33, 'Income'),
(44, 'Other')`

// FetchTransactionsTestCase Structure to hold each test case for FetchTransactions
type FetchTransactionsTestCase struct {
	Name       string
	DateFilter time.Time
	WantDesc   []string
	WantErr    bool
}

// FetchTransactionsTestCases Test cases for FetchTransactions
var FetchTransactionsTestCases = []FetchTransactionsTestCase{
	{
		"Fetch all transactions",
		time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
		[]string{"Groceries", "Rent", "Salary"},
		false,
	},
	{
		"Filter transactions after October 1st",
		time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC),
		[]string{"Rent", "Salary"},
		false,
	},
	{
		"Empty result - no transactions after October 10th",
		time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC),
		[]string{},
		false,
	},
	{
		"Invalid query - invalid date format",
		time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		[]string{},
		true,
	},
	{
		"Invalid query - invalid date value",
		time.Time{}, // Invalid zero date
		[]string{},
		true,
	},
}

// NewSqlite3HandlerTestCase Structure to hold each test case for NewSqlite3Handler
type NewSqlite3HandlerTestCase struct {
	Name    string
	DbFile  string
	WantErr bool
}

// NewSqlite3HandlerTestCases Test cases for NewSqlite3Handler
var NewSqlite3HandlerTestCases = []NewSqlite3HandlerTestCase{
	{
		"Valid in-memory database",
		":memory:",
		false,
	},
	{
		"Invalid database file path",
		"asd12.xx",
		true,
	},
}
