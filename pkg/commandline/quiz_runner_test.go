package commandline_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/millidavids/quiz/pkg/commandline"

	"github.com/millidavids/quiz/pkg"
)

func Test_QuizRunner(t *testing.T) {
	t.Run("Can run a quiz and get correct answer", should_get_correct_answer)
	t.Run("Can run q quiz and get incorrect answer", should_get_incorrect_answer)
}

func should_get_correct_answer(t *testing.T) {
	q := root.Quiz{Questions: map[string]int{"1+1": 2}}

	fake_stdin_runner(&q, "2")

	if q.Correct != 1 {
		t.Error("Did not get a correct answer")
	}
}

func should_get_incorrect_answer(t *testing.T) {
	q := root.Quiz{Questions: map[string]int{"1+1": 2}}

	if err := fake_stdin_runner(&q, "3"); err != nil {
		t.Fatal(err)
	}

	if q.Correct == 1 {
		t.Error("Got a correct answer")
	}
}

func fake_stdin_runner(q *root.Quiz, a string) error {
	qr := commandline.QuizRunner{}

	content := []byte(a)
	tmpfile, err := ioutil.TempFile("", "t")
	if err != nil {
		return err
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		return err
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		return err
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	os.Stdin = tmpfile

	qr.Run(q)

	return nil
}
