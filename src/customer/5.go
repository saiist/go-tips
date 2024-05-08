package customer

import "database/sql"

type ICustomerService interface {
	GetCustomerBalance(id int) (int, error)
}

type CustomerService struct {
	db *sql.DB
}

func NewCustomerService(db *sql.DB) ICustomerService {
	return &CustomerService{db: db}
}

func (cs *CustomerService) GetCustomerBalance(id int) (int, error) {
	var balance int
	err := cs.db.QueryRow("SELECT balance FROM customers WHERE id = ?", id).Scan(&balance)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

type IIntConfig interface {
	Get() int
}

type IntConfig struct {
}

type Foo struct {
	Threshold IIntConfig
}

func (c *IntConfig) Get() int {
	return 1
}

func (c *IntConfig) Set(i int) {
}

func NewFoo(threshold IIntConfig) Foo {
	return Foo{Threshold: threshold}
}
