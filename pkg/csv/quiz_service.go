package csv

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/millidavids/quiz/pkg"
)

// QuizService implements the root.QuizService
type QuizService struct {
	Filename  string
	TimeLimit int
}

// Create will generate a quiz from a csv file
func (qs *QuizService) Create() *root.Quiz {
	q := root.Quiz{
		Questions: make(map[string]int),
		TimeLimit: time.Duration(qs.TimeLimit) * time.Second,
	}
	f, _ := os.Open(qs.Filename)
	r := csv.NewReader(bufio.NewReader(f))

	for {
		l, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		q.Questions[l[0]], _ = strconv.Atoi(l[1])
	}

	return &q
}
