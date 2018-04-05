package commandline

import (
	"fmt"

	"github.com/millidavids/quiz/pkg"
)

type QuizRunner struct{}

func (qr *QuizRunner) Run(q *root.Quiz) {
	for k, v := range q.Questions {
		var a int
		fmt.Printf("%v = ", k)
		fmt.Scan(&a)
		if a == v {
			q.Correct++
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect!")
		}
	}
}
