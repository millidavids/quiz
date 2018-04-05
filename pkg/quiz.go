package root

type Quiz struct {
	Questions map[string]int
	Correct   int
}

type QuizService interface {
	Create() Quiz
}

type QuizRunner interface {
	Run(q *Quiz)
}
