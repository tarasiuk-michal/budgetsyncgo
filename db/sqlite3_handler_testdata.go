// Package db testdata.go
package db

// CreateTransactionsTableQuery Mock table creation query
const CreateTransactionsTableQuery = `
CREATE TABLE transactions (
	transaction_pk INTEGER PRIMARY KEY,
	name TEXT,
	amount REAL,
	category_fk INTEGER,
	date_created INTEGER
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
(1, 'Groceries', 50.0, 11, 1695168000), -- 2023-09-20
(2, 'Rent', 500.0, 22, 1696118400),	-- 2023-10-01
(3, 'Salary', 2000.0, 33, 1696550400), -- 2023-10-06
(4, 'Thing', 300.0, 44, 1601899200)	-- 2020-10-05
`

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
	DateFilter string
	WantDesc   []string
	WantErr    bool
}

// FetchTransactionsTestCases Test cases for FetchTransactions
var FetchTransactionsTestCases = []FetchTransactionsTestCase{
	{
		"Fetch all transactions",
		"2023-09-01",
		[]string{"Groceries", "Rent", "Salary"},
		false,
	},
	{
		"Filter transactions after October 1st",
		"2023-10-01",
		[]string{"Rent", "Salary"},
		false,
	},
	{
		"Empty result - no transactions after October 10th",
		"2023-10-10",
		[]string{},
		false,
	},
	{
		"Invalid query - invalid date format",
		"2020-01-01 12:00:00",
		[]string{},
		true,
	},
	{
		"Invalid query - invalid date value",
		"", // Invalid zero date
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
		"Valid database file path",
		"test_db.sqlite",
		false,
	},
	{
		"Invalid database file path",
		"asd12.xx",
		true,
	},
}
