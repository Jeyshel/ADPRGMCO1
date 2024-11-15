/*
Last names: Gaffud, Jean Luc G.
Language: Go
Paradigm(s): Imperative
*/

package main

import (
	f "fmt"
)

func main() {
	var loanAmount float64
	var annualRate float64
	var yearlyTerm float64

	f.Println("Enter loan amount in PHP: ")
	f.Scanln(&loanAmount)

	f.Println("Enter annual interest rate in percent: ")
	f.Scanln(&annualRate)

	f.Println("Enter loan yearly term: ")
	f.Scanln(&yearlyTerm)

	annualRatePercent := annualRate / 100

	monthlyRate := annualRatePercent / 12
	monthlyTerm := yearlyTerm * 12
	totalInterest := loanAmount * monthlyRate * monthlyTerm
	monthlyPayment := (loanAmount + totalInterest) / monthlyTerm

	f.Println("Loan Amount: PHP ", loanAmount)
	f.Println("Annual Interest Rate: ", annualRate, "%")
	f.Println("Loan Term: ", monthlyTerm, " months")
	f.Printf("Monthly Repayment: PHP %.2f\n", monthlyPayment)
	f.Printf("Total Interest: PHP %.2f\n", totalInterest)
}
