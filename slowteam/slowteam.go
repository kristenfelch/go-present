package slowteam

import (
	"fmt"
	"github.com/kristenfelch/go-present/student"
	"math/rand"
	"time"
)

/*
 One of the prime benefits to functional programming is the performance enhancements.
 Let's use a function that actual takes awhile to run to calculate student GPAs.
*/

/////////////////////////////////// HigherOrderTeam //////////////////////////////////////////

// HigherOrderTeam is meant to be a copy of the best we could do prior, but applied to only
// 10 students with long-running calculation functions.

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
			student.Map(make([]string, 5),
				func(string) int {
					return calculateHigherOrderStudentGPA(skills)
				}),
			func(old, new int) int {
				return old + new
			}),
		5)
	//abstract out average calculation
	fmt.Printf("Final Team Average: %d\n", teamAverage)
}

//chooses skills, gets grades, and calculates student GPA
func calculateHigherOrderStudentGPA(skills []string) int {
	time.Sleep(1 * time.Second)
	numberOfSkills := rand.Intn(4) + 1
	skillsLearned := skills[0:numberOfSkills]
	return student.LearnGolang(skillsLearned)
}

func calculateAverage(sum, length int) int {
	return sum / length
}

///////////////////////////////////////// WinningTeam //////////////////////////////////////////////

type WinningTeam struct{}

func (WinningTeam) Run() {
	channel := make(chan int)
	count := 0
	student.Map(make([]string, 5),
		//should rewrite Map to return nothing, so we don't have to fake a 0 return
		func(string) int {
			go calculateWinningStudentGPA(channel, skills)
			return 0
		})
	student.Map(make([]string, 5),
		func(string) int {
			count += <-channel
			return 0
		})
	//abstract out average calculation
	fmt.Printf("Final Team Average: %d\n", calculateAverage(count, 5))
}

//chooses skills, gets grades, and calculates student GPA
func calculateWinningStudentGPA(channel chan int, skills []string) {
	time.Sleep(1 * time.Second)
	numberOfSkills := rand.Intn(4) + 1
	skillsLearned := skills[0:numberOfSkills]
	channel <- student.LearnGolang(skillsLearned)
}

// Best solution is a combination of :
// Extracting Pure Functions with no side effects
// Utilizing higher order functions (map) to apply functions to many inputs or data points
// Using goroutines to obtain parallel processing, which is only possible for pure functions.
