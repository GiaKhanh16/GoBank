package main

import (
	"database/sql"

	"fmt"

	_ "github.com/lib/pq"
)



type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	GetAccounts() ([]*Account, error)
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func (s *PostgresStorage) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStorage) createAccountTable() error {
	query := `create table if not exists accounts (
    id serial primary key,
    first_name varchar(50),
    last_name varchar(50),
    number serial,
    balance serial,
    created_at timestamp
)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStorage) CreateAccount(acc *Account) error {

	query := `INSERT INTO accounts (first_name, last_name, number,
	balance, created_at) values ($1, $2, $3, $4, $5)`

	resp, err := s.db.Exec(query, acc.FirstName, acc.LastName, acc.Number, acc.Balance, acc.CreatedAt)

	if err != nil {
		return err
	}
	fmt.Println(resp)

	return nil
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}



func (s *PostgresStorage) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStorage) DeleteAccount(id int) error {

	_, err := s.db.Exec("delete from accounts where id = $1", id)

	return err
}

func (s *PostgresStorage) GetAccountByID(id int) (*Account, error) {
	rows, err := s.db.Query("select * from accounts where id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return ScanIntoAccount(rows)
	}

	return nil, nil
}
func (s *PostgresStorage) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("SELECT * FROM accounts")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := ScanIntoAccount(rows)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func ScanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	)
	return account, err
}
