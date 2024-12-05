package main

import "fmt"

// PaymentMethod interface defines the core behavior for payments
type PaymentMethod interface {
	Pay(amount float64)
}

// Refundable interface extends PaymentMethod to add refund functionality
type Refundable interface {
	PaymentMethod
	Refund(amount float64)
}

// CreditCard struct represents a credit card payment
type CreditCard struct {
	CardNumber string
}

func (cc CreditCard) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using Credit Card ending in %s.\n", amount, cc.CardNumber[len(cc.CardNumber)-4:])
}
// 2.4653-> 2.46

func (cc CreditCard) Refund(amount float64) {
	fmt.Printf("Refunded $%.2f to Credit Card ending in %s.\n", amount, cc.CardNumber[len(cc.CardNumber)-4:])
}

// PayPal struct represents a PayPal payment
type PayPal struct {
	Email string
}

func (pp PayPal) Pay(amount float64) {
	fmt.Printf("Paid $%.2f using PayPal account: %s.\n", amount, pp.Email)
}

func (pp PayPal) Refund(amount float64) {
	fmt.Printf("Refunded $%.2f to PayPal account: %s.\n", amount, pp.Email)
}

// Cash struct represents a cash payment
type Cash struct{}

func (c Cash) Pay(amount float64) {
	fmt.Printf("Paid $%.2f in cash.\n", amount)
}

// ProcessPayment processes any type that implements the PaymentMethod interface
func ProcessPayment(pm PaymentMethod, amount float64) {
	pm.Pay(amount)
}

// ProcessRefund processes any type that implements the Refundable interface
func ProcessRefund(r Refundable, amount float64) {
	r.Refund(amount)
}

func main() {
	// Instances of different payment methods
	creditCard := CreditCard{CardNumber: "1234567812345678"}
	payPal := PayPal{Email: "user@example.com"}
	cash := Cash{}

	// Processing payments
	fmt.Println("Processing Payments:")
	ProcessPayment(creditCard, 100.50)
	ProcessPayment(payPal, 50.75)
	ProcessPayment(cash, 20.00)

	fmt.Println("\nProcessing Refunds:")
	// Processing refunds for refundable methods
	ProcessRefund(creditCard, 30.25)
	ProcessRefund(payPal, 10.50)

	// Cash does not implement Refundable, so it cannot be refunded
	// Uncommenting the next line would cause a compile-time error
	// ProcessRefund(cash, 5.00)
}
