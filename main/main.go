package main

import (
	"github.com/kristenfelch/go-present/demo"
	"github.com/kristenfelch/go-present/monitor"
	//"github.com/kristenfelch/go-present/slowteam"
	"github.com/kristenfelch/go-present/team"
)

/*
  Calculate team average for our Golang students based on several different paradigms.
*/
func main() {
	// Utilizing an OOP system - mutating data, for loops, no higher order functions.
	demo.Decorate(team.OOPTeam{}, monitor.Monitor("OOP Team Average")).Run()

	// Abstracting out some preliminary logic into functions
	//demo.Decorate(team.FunctioningTeam{}, monitor.Monitor("Functioning Team Average")).Run()

	// Recursion rather than iteration to maintain immutable data
	//demo.Decorate(team.RecursiveTeam{}, monitor.Monitor("Recursive Team Average")).Run()

	// Tailwise Recursion in hopes of solving stack overflow issue
	//demo.Decorate(team.TailingTeam{}, monitor.Monitor("Tail Recursion Team Average")).Run()

	// Using channels to mimic tailwise recursion
	//demo.Decorate(team.ChannelingTeam{}, monitor.Monitor("Channeling Team Average")).Run()

	// Adding higher order functions to our initial functioning team
	//demo.Decorate(team.HigherOrderTeam{}, monitor.Monitor("Higher Order Team Average")).Run()

	//demo.Decorate(slowteam.HigherOrderTeam{}, monitor.Monitor("Slow Higher Order Team Average")).Run()

	//demo.Decorate(slowteam.WinningTeam{}, monitor.Monitor("Slow Winning Team Average")).Run()

}
