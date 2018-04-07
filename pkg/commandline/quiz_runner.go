package commandline

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/millidavids/quiz/pkg"
)

type QuizRunner struct{}

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
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect!")
		}
	}
}
