package root

import "time"

// Quiz type contains a set of Questions in map form and keeps track of Correct
// answers and a quiz TimeLimit
type Quiz struct {
	Questions map[string]int
	Correct   int
	TimeLimit time.Duration
}

// QuizService interface defines functionality for creating quizzes
type QuizService interface {
	Create() Quiz
}

// QuizRunner interface defiles functionality for running quizzes
type QuizRunner interface {
	Run(q *Quiz)
	TimedRun(q *Quiz)
}
