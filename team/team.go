package team

import (
	"fmt"
	"github.com/kristenfelch/go-present/student"
	"math/rand"
)

/*
 Team calculates GPA of all students learning GoLang and
 prints out the overall team average.
 There are for this use case, 10M students learning Golang (For the purposes of performance analysis)
*/

/////////////////////////////////////////// OOPTeam///////////////////////////////////////////

type OOPTeam struct{}

func (OOPTeam) Run() {
	teamTotal := 0
	//skills that students may opt to learn, slice.
	skills := []string{
		"Channels",
		"Types",
		"Pointers",
		"Structs",
		"Slices"}
	//as 10M students to learn Golang
	for i := 0; i < 10000000; i++ {
		numberOfSkills := rand.Intn(4) + 1
		skillsCompleted := skills[0:numberOfSkills]
		student := student.OOPsStudent{Skills: skillsCompleted}
		teamTotal = teamTotal + student.LearnGolang()
	}
	fmt.Printf("OOP Team Average: %d\n", teamTotal/10000000)
}

//Again, we see the OOP paradigm of attaching the method to the object, we see instantiation of
//new objects for each student in our 10M.

//////////////////////////////////////// FunctioningTeam //////////////////////////////////////

// Computation as the evaluation of mathematical functions.
// Functions are accessible throughout program, not attached to objects.

type FunctioningTeam struct{}

func (FunctioningTeam) Run() {
	skills := []string{
		"Channels",
		"Types",
		"Pointers",
		"Structs",
		"Slices"}
	teamTotal := 0
	for i := 0; i < 10000000; i++ {
		//abstract out functionality for determining skills studied for each student
		studentGPA := calculateStudentGPA(skills)
		teamTotal = teamTotal + studentGPA
	}
	//abstract out average calculation
	fmt.Printf("Functioning Team Average: %d\n", calculateAverage(teamTotal, 10000000))
}

//chooses skills, gets grades, and calculates student GPA
func calculateStudentGPA(skills []string) int {
	numberOfSkills := rand.Intn(4) + 1
	skillsLearned := skills[0:numberOfSkills]
	student := student.FunctioningStudent{Skills: skillsLearned}
	return student.LearnGolang()
}

func calculateAverage(sum, length int) int {
	return sum / length
}

//We really didn't change much.  In fact, the execution flow, memory usage, and performance should be
//pretty much the same.  But we started honing in on a new pattern of focusing on functions rather
//than objects.

//////////////////////////////////////// RecursiveTeam ///////////////////////////////////////

// Immutable State.
// Variable values cannot be changed.
// Iteration cannot be used since by definition we have at least a counter.

type RecursiveTeam struct{}

func (RecursiveTeam) Run() {
	skills := []string{
		"Channels",
		"Types",
		"Pointers",
		"Structs",
		"Slices"}
	teamTotal := recursivelyCalculateTeamTotal(0, skills)
	//abstract out average calculation
	fmt.Printf("Recursive Team Average: %d\n", calculateAverage(teamTotal, 10000000))
}

func recursivelyCalculateTeamTotal(index int, skills []string) int {
	if index >= 10000000 {
		return 0
	}
	return calculateStudentGPARecursive(skills) + recursivelyCalculateTeamTotal(index+1, skills)
}

func calculateStudentGPARecursive(skills []string) int {
	numberOfSkills := rand.Intn(4) + 1
	skillsLearned := skills[0:numberOfSkills]
	//only difference here is type of student
	student := student.RecursiveStudent{Skills: skillsLearned}
	return student.LearnGolang()
}

// Advantages? Typically less code.  Functionality may appear more elegant and easier to read.
// For some algorithms, being able to recurse in multiple places within the function are possible.
// However, we obviously can't do this because of our stack issue!

//////////////////////////////////////// TailingTeam //////[[/////////////////////////////////

// Replace recursion with tail recursion in hopes of avoiding stack overflow.

type TailingTeam struct{}

func (TailingTeam) Run() {
	skills := []string{
		"Channels",
		"Types",
		"Pointers",
		"Structs",
		"Slices"}
	teamTotal := tailCalculateTeamTotal(0, 0, skills)
	//abstract out average calculation
	fmt.Printf("Final Team Average: %d\n", calculateAverage(teamTotal, 10000000))
}

func tailCalculateTeamTotal(index, runningTotal int, skills []string) int {
	if index >= 10000000 {
		return runningTotal
	}
	studentGPA := calculateStudentGPATailwise(skills)
	return tailCalculateTeamTotal(index+1, runningTotal+studentGPA, skills)
}

// only difference is we are using a TailStudent
func calculateStudentGPATailwise(skills []string) int {
	numberOfSkills := rand.Intn(4) + 1
	skillsLearned := skills[0:numberOfSkills]
	//only difference here is type of student
	student := student.TailStudent{Skills: skillsLearned}
	return student.LearnGolang()
}

//Tail recursion optimization is not supported in Golang. So... is there any way around this??

//////////////////////////////////////// ChannelingTeam ///////////////////////////////////////

// Let's imitate tail recursion by using a channel!  Goroutines can be spawned  for each new
// recursive step, and the final result is awaited by waiting on a result placed in a channel.
type ChannelingTeam struct{}

func (ChannelingTeam) Run() {
	skills := []string{
		"Channels",
		"Types",
		"Pointers",
		"Structs",
		"Slices"}
	result := make(chan int)
	channelCalculateTeamTotal(0, 0, result, skills)
	teamTotal := <-result
	//abstract out average calculation
	fmt.Printf("Channeling Team Average: %d\n", calculateAverage(teamTotal, 10000000))
}

func channelCalculateTeamTotal(index, runningTotal int, result chan int, skills []string) {
	if index >= 10000000 {
		result <- runningTotal
	}
	studentGPA := calculateStudentGPAChanneling(skills)
	go channelCalculateTeamTotal(index+1, runningTotal+studentGPA, result, skills)
}

// only difference is we are using a ChannelingStudent
func calculateStudentGPAChanneling(skills []string) int {
	numberOfSkills := rand.Intn(4) + 1
	skillsLearned := skills[0:numberOfSkills]
	//only difference here is type of student
	student := student.ChannelingStudent{Skills: skillsLearned}
	return student.LearnGolang()
}

// OK so we CAN make this functional.  But, is it worth it in terms of performance?
// Spawning routines has the cost of scheduling built in.

/////////////////////////////////// HigherOrderTeam //////////////////////////////////////////

// Let's get down to focusing on functions - as arguments and return results.  This will allow
// us to forget our object abstraction and focus on behavior encapsulated into functions.
// Higher order, pure functions are the goal.

type HigherOrderTeam struct{}

var skills = []string{
	"Channels",
	"Types",
	"Pointers",
	"Structs",
	"Slices"}

func (HigherOrderTeam) Run() {

	teamAverage := calculateAverage(
		//note ability to abstract out and reuse functions.
		student.Reduce(
			student.Map(make([]string, 10000000),
				func(string) int {
					return calculateHigherOrderStudentGPA(skills)
				}),
			func(old, new int) int {
				return old + new
			}),
		10000000)
	//abstract out average calculation
	fmt.Printf("Final Team Average: %d\n", teamAverage)
}

//chooses skills, gets grades, and calculates student GPA
func calculateHigherOrderStudentGPA(skills []string) int {
	numberOfSkills := rand.Intn(4) + 1
	skillsLearned := skills[0:numberOfSkills]
	return student.LearnGolang(skillsLearned)
}
