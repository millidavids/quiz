package commandline_test

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/millidavids/quiz/pkg/commandline"

	"github.com/millidavids/quiz/pkg"
)

func Test_QuizRunner(t *testing.T) {
	t.Run("Can run a quiz and get correct answer", should_get_correct_answer)
	t.Run("Can run a quiz and get incorrect answer", should_get_incorrect_answer)
	t.Run("Can run a quiz and timeout", should_timeout_with_error)
}

func should_get_correct_answer(t *testing.T) {
	q := root.Quiz{
		Questions: map[string]int{"1+1": 2},
		TimeLimit: time.Duration(100 * time.Hour),
	}
	qr := commandline.QuizRunner{}

	fake_stdin_runner(&q, "2", qr.Run)

	if q.Correct != 1 {
		t.Error("did not get a correct answer")
	}
}

func should_get_incorrect_answer(t *testing.T) {
	q := root.Quiz{
		Questions: map[string]int{"1+1": 2},
		TimeLimit: time.Duration(100 * time.Hour),
	}
	qr := commandline.QuizRunner{}

	if err := fake_stdin_runner(&q, "3", qr.Run); err != nil {
		t.Fatal(err)
	}

	if q.Correct == 1 {
		t.Error("got a correct answer")
	}
}

func should_timeout_with_error(t *testing.T) {
	q := root.Quiz{
		Questions: map[string]int{
			"1+1": 2,
			"1+2": 3,
			"1+3": 4,
			"1+4": 5,
			"1+5": 6,
			"1+6": 7,
			"1+7": 8,
		},
		TimeLimit: time.Duration(0),
	}

	qr := commandline.QuizRunner{}

	if err := fake_stdin_runner(&q, "\n2\n3\n4\n5\n6\n7\n8", qr.TimedRun); err == nil {
		t.Error("did not timeout")
	}
}

func fake_stdin_runner(q *root.Quiz, a string, f func(*root.Quiz) error) error {
	c := []byte(a)
	tf, err := ioutil.TempFile("", "t")
	if err != nil {
		return err
	}

	defer os.Remove(tf.Name())

	if _, err := tf.Write(c); err != nil {
		return err
	}

	if _, err := tf.Seek(0, 0); err != nil {
		return err
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	os.Stdin = tf

	return f(q)
}
