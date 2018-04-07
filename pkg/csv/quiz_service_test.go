package csv_test

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/millidavids/quiz/pkg/csv"
)

func Test_QuizService(t *testing.T) {
	t.Run("Can create a Quiz from a csv file", should_create_quiz_from_csv_file)
}

func should_create_quiz_from_csv_file(t *testing.T) {
	content := []byte("1+1,2")
	tmpfile, err := ioutil.TempFile("", "t")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}

	qs := csv.QuizService{Filename: tmpfile.Name()}
	q := qs.Create()

	if typ := reflect.TypeOf(q).String(); typ != "*root.Quiz" {
		t.Errorf("Created the wrong object type: %v", typ)
	}

	if qqs := q.Questions; qqs["1+1"] != 2 {
		t.Errorf("Created the map of questions incorrectly: %v", qqs)
	}
}
