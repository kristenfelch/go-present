package main

import (
	"github.com/kristenfelch/go-present/demo"
	"github.com/kristenfelch/go-present/monitor"
	"github.com/kristenfelch/go-present/reportcard"
)

/*
  Calculate class averages based on several different paradigms.
*/
func main() {
	// Utilizing an OOP system - mutating data, for loops, no higher order functions.
	demo.Decorate(reportcard.OOPReportCard{}, monitor.Monitor("OOP Class Average")).Run()

	// Abstracting out some preliminary logic into functions
	demo.Decorate(reportcard.FunctioningReportCard{}, monitor.Monitor("Functioning Class Average")).Run()

	// Recursion rather than iteration to maintain immutable data
	//demo.Decorate(reportcard.RecursiveReportCard{}, monitor.Monitor("Recursive Class Average")).Run()

	// Tailwise Recursion in hopes of solving stack overflow issue
	//demo.Decorate(reportcard.TailingReportCard{}, monitor.Monitor("Tail Recursion Class Average")).Run()

	// Using channels to mimic tailwise recursion
	demo.Decorate(reportcard.ChannelingReportCard{}, monitor.Monitor("Channeling for Recursion (Abbreviated)")).Run()

	demo.Decorate(reportcard.HigherOrderReportCard{}, monitor.Monitor("Higher Order Functions")).Run()
}
