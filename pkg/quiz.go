package root

// Quiz type contains a set of Questions in map form and keeps track of Correct answers
type Quiz struct {
	Questions map[string]int
	Correct   int
}

// QuizService interface defines functionality for creating quizzes
type QuizService interface {
	Create() Quiz
}

// QuizRunner interface defiles functionality for running quizzes
type QuizRunner interface {
	Run(q *Quiz)
}
