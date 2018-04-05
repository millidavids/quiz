package csv

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/millidavids/quiz/pkg"
)

type QuizService struct {
	Filename string
}

func NewQuizService(f string) *QuizService {
	return &QuizService{Filename: f}
}

func (qs *QuizService) Create() *root.Quiz {
	q := root.Quiz{Questions: make(map[string]int)}
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
