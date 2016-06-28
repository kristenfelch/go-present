package reportcard

import (
	"fmt"
	"github.com/kristenfelch/go-present/student"
	"math/rand"
)

/*
 ReportCard calculates GPA of all students learning GoLang and
 prints out the overall class/course average.
 There are for this use case, 1M students learning Golang.
*/

/////////////////////////////////////////// OOPReportCard///////////////////////////////////////////

type OOPReportCard struct{}

func (OOPReportCard) Run() {
	classTotal := 0
	//skills that students may opt to learn.
	skills := []string{
		"Channels",
		"Types",
		"Pointers",
		"Structs",
		"Slices"}
	for i := 0; i < 10000000; i++ {
		numberOfCourses := rand.Intn(4) + 1
		coursesCompleted := skills[0:numberOfCourses]
		student := student.OOPsStudent{Concepts: coursesCompleted}
		classTotal = classTotal + student.LearnGolang()
	}
	fmt.Printf("Final Class Average: %d\n", classTotal/1000000)
}

//////////////////////////////////////// FunctioningReportCard //////////////////////////////////////

// Computation as the evaluation of mathematical functions.
// Functions are accessible throughout program, not attached to objects.

type FunctioningReportCard struct{}

func (FunctioningReportCard) Run() {
	skills := []string{
		"Channels",
		"Types",
		"Pointers",
		"Structs",
		"Slices"}
	classTotal := 0
	for i := 0; i < 10000000; i++ {
		//abstract out functionality for determining skills studied for each student
		studentGPA := calculateStudentGPA(skills)
		classTotal = classTotal + studentGPA
	}
	//abstract out average calculation
	fmt.Printf("Final Class Average: %d\n", calculateAverage(classTotal, 1000000))
}

//chooses skills, gets grades, and calculates student GPA
func calculateStudentGPA(skills []string) int {
	numberOfSkills := rand.Intn(4) + 1
	skillsLearned := skills[0:numberOfSkills]
	student := student.FunctioningStudent{Concepts: skillsLearned}
	return student.LearnGolang()
}

func calculateAverage(sum, length int) int {
	return sum / length
}

//We really didn't change much.  In fact, the execution flow, memory usage, and performance should be
//pretty much the same.  But we started honing in on a new pattern of focusing on functions rather
//than objects.

//////////////////////////////////////// RecursiveReportCard ///////////////////////////////////////

// Immutable State.
// Variable values cannot be changed.
// Iteration cannot be used since by definition we have at least a counter.

type RecursiveReportCard struct{}

func (RecursiveReportCard) Run() {
	skills := []string{
		"Channels",
		"Types",
		"Pointers",
		"Structs",
		"Slices"}
	classTotal := recursivelyCalculateClassTotal(0, skills)
	//abstract out average calculation
	fmt.Printf("Final Class Average: %d\n", calculateAverage(classTotal, 1000000))
}

func recursivelyCalculateClassTotal(index int, skills []string) int {
	if index >= 10000000 {
		return 0
	}
	return calculateStudentGPARecursive(skills) + recursivelyCalculateClassTotal(index+1, skills)
}

func calculateStudentGPARecursive(skills []string) int {
	numberOfSkills := rand.Intn(4) + 1
	skillsLearned := skills[0:numberOfSkills]
	//only difference here is type of student
	student := student.RecursiveStudent{Concepts: skillsLearned}
	return student.LearnGolang()
}

// Advantages? Typically less code.  Functionality may appear more elegant and easier to read.
// For some algorithms, being able to recurse in multiple places within the function are possible.
// However, we obviously can't do this because of our stack issue!

//////////////////////////////////////// TailingReportCard //////[[/////////////////////////////////

// Replace recursion with tail recursion in hopes of avoiding stack overflow.

type TailingReportCard struct{}

func (TailingReportCard) Run() {
	skills := []string{
		"Channels",
		"Types",
		"Pointers",
		"Structs",
		"Slices"}
	classTotal := tailCalculateClassTotal(0, 0, skills)
	//abstract out average calculation
	fmt.Printf("Final Class Average: %d\n", calculateAverage(classTotal, 1000000))
}

func tailCalculateClassTotal(index, runningTotal int, skills []string) int {
	if index >= 10000000 {
		return runningTotal
	}
	studentGPA := calculateStudentGPATailwise(skills)
	return tailCalculateClassTotal(index+1, runningTotal+studentGPA, skills)
}

// only difference is we are using a TailStudent
func calculateStudentGPATailwise(skills []string) int {
	numberOfSkills := rand.Intn(4) + 1
	skillsLearned := skills[0:numberOfSkills]
	//only difference here is type of student
	student := student.TailStudent{Concepts: skillsLearned}
	return student.LearnGolang()
}

//Tail recursion optimization is not supported in Golang. So... is there any way around this??

//////////////////////////////////////// ChannelingReportCard ///////////////////////////////////////

// Let's imitate tail recursion by using a channel!  Goroutines can be spawned  for each new
// recursive step, and the final result is awaited by waiting on a result placed in a channel.
type ChannelingReportCard struct{}

func (ChannelingReportCard) Run() {
	skills := []string{
		"Channels",
		"Types",
		"Pointers",
		"Structs",
		"Slices"}
	result := make(chan int)
	channelCalculateClassTotal(0, 0, result, skills)
	classTotal := <-result
	//abstract out average calculation
	fmt.Printf("Final Class Average: %d\n", calculateAverage(classTotal, 1000000))
}

func channelCalculateClassTotal(index, runningTotal int, result chan int, skills []string) {
	if index >= 1000000 {
		result <- runningTotal
	}
	studentGPA := calculateStudentGPAChanneling(skills)
	go channelCalculateClassTotal(index+1, runningTotal+studentGPA, result, skills)
}

// only difference is we are using a ChannelingStudent
func calculateStudentGPAChanneling(skills []string) int {
	numberOfSkills := rand.Intn(4) + 1
	skillsLearned := skills[0:numberOfSkills]
	//only difference here is type of student
	student := student.ChannelingStudent{Concepts: skillsLearned}
	return student.LearnGolang()
}

// OK so we CAN make this functional.  But, no it's not worth it in terms of performance.
// This is where we would have to choose practicality over purity.

/////////////////////////////////// HigherOrderReportCard //////////////////////////////////////////

// Let's get down to focusing on functions - as arguments and return results.  This will allow
// us to forget our object abstraction and focus on behavior encapsulated into functions.
// Higher order, pure functions are the goal.

type HigherOrderReportCard struct{}

func (HigherOrderReportCard) Run() {
	skills := []string{
		"Channels",
		"Types",
		"Pointers",
		"Structs",
		"Slices"}

	classAverage := calculateAverage(
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
	fmt.Printf("Final Class Average: %d\n", classAverage)
}

//chooses skills, gets grades, and calculates student GPA
func calculateHigherOrderStudentGPA(skills []string) int {
	numberOfSkills := rand.Intn(4) + 1
	skillsLearned := skills[0:numberOfSkills]
	return student.LearnGolang(skillsLearned)
}
