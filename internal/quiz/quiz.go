package quiz

type Strategy func(ans Answer) bool

type Quiz struct {
	id      string
	Title   string `json:"title"`
	Problem string `json:"problem"`
}

func (quiz Quiz) createId() {
	quiz.id = quiz.Title + quiz.Problem
}

func (quiz Quiz) save() {
	quiz.createId()
}

type Answer interface {
}
