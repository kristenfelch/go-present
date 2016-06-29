package student

import (
	"math/rand"
)

////////////////////////////////////////// OOPStudent //////////////////////////////////////////////
type OOPsStudent struct {
	//Grade point average of all Golang skills
	Average int
	//Skills that this student learned.
	Skills []string
}

func (s *OOPsStudent) LearnGolang() int {
	sum := 0

	//obtain grade for each skill
	for _ = range s.Skills {
		grade := rand.Intn(50) + 50
		sum = sum + grade
	}
	//calculate GPA for student
	s.Average = sum / len(s.Skills)
	return s.Average
}

//Object oriented paradigm of attaching methods and attributes/fields to an object,
//capturing the data and the functionality in the same place!

///////////////////////////////////////// FunctioningStudent ///////////////////////////////////////

// Computation as the evaluation of mathematical functions.
// Functions are accessible throughout program, not attached to objects.

type FunctioningStudent struct {
	Average int
	Skills  []string
}

func (s *FunctioningStudent) LearnGolang() int {
	studentTotal := calculateTotal(s.Skills)
	return calculateAverage(studentTotal, len(s.Skills))
}

func calculateTotal(skills []string) int {
	sum := 0

	//obtain grade for each skill
	for _ = range skills {
		grade := rand.Intn(50) + 50
		sum = sum + grade
	}
	return sum
}

func calculateAverage(total int, number int) int {
	return total / number
}

///////////////////////////////////////// RecursiveStudent /////////////////////////////////////////

// Iteration is a characteristic of procedural programming.  Recursion is a feature of functional
// programming, as no mutable counters or data objects are required.

type RecursiveStudent struct {
	Average int
	Skills  []string
}

func (s *RecursiveStudent) LearnGolang() int {
	studentTotal := calculateTotalRecursively(0, s.Skills)
	return calculateAverage(studentTotal, len(s.Skills))
}

//calculates a grade for each skill, adding this grade to sum of total for remainder of slice.
func calculateTotalRecursively(index int, skills []string) int {
	if index >= len(skills) {
		return 0
	}
	return rand.Intn(50) + 50 + calculateTotalRecursively(index+1, skills)
}

////////////////////////////////////////// TailStudent /////////////////////////////////////////////

// Tail recursion is a particular way of writing recursive algorithms so that we avoid stack overflow.
// Basically, our final return value of the recursive function _is_ the recursive call, so that
// a stack element can be replaced rather than added to.

type TailStudent struct {
	Average int
	Skills  []string
}

func (s *TailStudent) LearnGolang() int {
	studentTotal := calculateTotalTailwise(0, 0, s.Skills)
	return calculateAverage(studentTotal, len(s.Skills))
}

//calculates a grade for each skill, passing the new index _and_ runningTotal to next interation.
func calculateTotalTailwise(index int, runningTotal int, skills []string) int {
	if index >= len(skills) {
		return runningTotal
	}
	score := rand.Intn(50) + 50
	return calculateTotalTailwise(index+1, runningTotal+score, skills)
}

//////////////////////////////////////// ChannelingStudent /////////////////////////////////////////

// Let's imitate tail recursion by using a channel!  Goroutines can be spawned  for each new
// recursive step, and the final result is awaited by waiting on a result placed in a channel.

type ChannelingStudent struct {
	Average int
	Skills  []string
}

func (s *ChannelingStudent) LearnGolang() int {
	result := make(chan int)
	calculateTotalChanneling(0, 0, result, s.Skills)
	studentTotal := <-result
	return calculateAverage(studentTotal, len(s.Skills))
}

//calculates a grade for each skill, passing new total through a channel.
func calculateTotalChanneling(index, runningTotal int, result chan int, skills []string) {
	if index >= len(skills) {
		result <- runningTotal
	}
	score := rand.Intn(50) + 50
	go calculateTotalChanneling(index+1, runningTotal+score, result, skills)
}

///////////////////////////////////////// HigherOrderStudent ///////////////////////////////////////

// Lose objects completely.
// Expressions rather than statements.
// Focus on extracting behavior to functions alone.
// Higher order functions are passed into and returned from other functions.

func LearnGolang(skills []string) int {
	return calculateAverage(
		calculateTotalHigherOrder(skills), len(skills))
}

func calculateTotalHigherOrder(skills []string) int {
	return Reduce(
		Map(skills, func(string) int {
			return getRandomGrade()
		}), func(oldValue, newValue int) int {
			return oldValue + newValue
		})
}

func getRandomGrade() int {
	return rand.Intn(50) + 50
}

//Map and Reduce are higher order functions - they take in other functions as arguments.
func Map(vs []string, f func(string) int) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func Reduce(vs []int, f func(int, int) int) int {
	total := 0
	for _, v := range vs {
		total = f(total, v)
	}
	return total
}

//Well, we're back to working!  No stack overflow.  About twice the time of our OO version.
// Why is this??
