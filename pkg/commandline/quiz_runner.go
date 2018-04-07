package commandline

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/millidavids/quiz/pkg"
)

// QuizRunner implements the root.QuizRunner interface
type QuizRunner struct{}

// Run will run the quiz questions through the command line
func (qr *QuizRunner) Run(q *root.Quiz) {
	s := bufio.NewScanner(os.Stdin)
	for k, v := range q.Questions {
		var a string
		vs := strconv.Itoa(v)
		fmt.Printf("%v = ", k)
		if s.Scan() {
			a = s.Text()
		}
		if a == vs {
			q.Correct++
		}
	}
}
