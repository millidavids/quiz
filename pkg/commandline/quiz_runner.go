package commandline

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/millidavids/quiz/pkg"
)

// QuizRunner implements the root.QuizRunner interface
type QuizRunner struct{}

// Run will run the quiz questions through the command line
func (qr *QuizRunner) Run(q *root.Quiz) error {
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
	return nil
}

// TimedRun will run the quiz with a timer and exit early when it expires
func (qr *QuizRunner) TimedRun(q *root.Quiz) error {
	fmt.Print("Press 'Enter' to start the quiz...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	eChan := make(chan error)

	t := time.NewTimer(q.TimeLimit)
	go func() {
		<-t.C
		eChan <- fmt.Errorf("\nYou ran out of time")
	}()
	go func() {
		qr.Run(q)
		eChan <- nil
	}()
	return <-eChan
}
