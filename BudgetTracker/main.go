package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Transaction struct to hold transaction details
type Transaction struct {
	ID       int
	Amount   float64
	Category string
	Date     time.Time
	Type     string
}

// budget tracker struct to manage transaction
type BudgetTracker struct {
	transactions []Transaction
	nextId       int
}

type FinancialRecord interface {
	GetAmount() float64
	GetType() string
}

func (t Transaction) GetAmount() float64 {
	return t.Amount
}
func (t Transaction) GetType() string {
	return t.Type
}
func (bg *BudgetTracker) AddTransaction(amount float64, category, ttype string) {
	newTransaction := Transaction{
		ID:       bg.nextId,
		Amount:   amount,
		Category: category,
		Date:     time.Now(),
		Type:     ttype,
	}
	bg.transactions = append(bg.transactions, newTransaction)
	bg.nextId++
}

func (bg BudgetTracker) Displaytransactions() {
	fmt.Println("ID\tAmount\tCategory\tDate\tType\t")
	for _, transaction := range bg.transactions {
		fmt.Printf("%d\t%2f\t%s\t%s\t%s\n", transaction.ID, transaction.Amount, transaction.Category, transaction.Date.Format("2000-01-02"), transaction.Type)
	}
}
func (bg BudgetTracker) getTotal(tType string) float64 {
	var total float64
	for _, transaction := range bg.transactions {
		if transaction.Type == tType {
			total += transaction.Amount
		}
	}
	return total
}

// save transaction to csv file
func (bg BudgetTracker) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)                                      //creating new csv file
	defer writer.Flush()                                               //it esures that data is writen to file before file closes
	writer.Write([]string{"ID", "Amount", "Category", "Date", "Type"}) //write to csv header

	//write data
	for _, t := range bg.transactions {
		record := []string{
			strconv.Itoa(t.ID),
			fmt.Sprintf("%.2f", t.Amount),
			t.Category,
			t.Date.Format("2000-01-02"),
			t.Type,
		}
		writer.Write(record)
	}
	fmt.Println("transaction saved to", file)
	return nil
}
func main() {
	//AddTransaction()
}
