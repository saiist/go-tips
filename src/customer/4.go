package customer

type customer struct {
	balance int
}

func NewCustomer() *customer {
	return &customer{}
}

func (c *customer) Balance() int {
	return c.balance
}

func (c *customer) SetBalance(balance int) {
	c.balance = balance
}
