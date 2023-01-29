package factory

import (
	"log"
	"math/rand"
	"time"

	"github.com/zakoemon/diatonictool/internal/data"
	"github.com/zakoemon/diatonictool/internal/quiz"
)

type QuizType int

const (
	AnswerNthNoteOfMajor = iota
	AnswerNthNoteOfMinor = iota
)

const ()

func GetQuiz(quizType QuizType) quiz.Quiz {
	switch quizType {
	case AnswerNthNoteOfMajor:
		key := randomKeyGenerator(data.Major)
		return quiz.Quiz{Title: "quiz", Problem: key().Name}
	}

	log.Fatal()
	return quiz.Quiz{}
}

func genRand(num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(num)
}
