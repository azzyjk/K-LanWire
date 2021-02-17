package src

type questionChecker struct {
	checker *aiChecker
}

func newQuestionChecker() *questionChecker {

	newQuestionChecker := &questionChecker{}
	newQuestionChecker.checker = newAiChecker("question", "question_check_mq_val.txt")
	return newQuestionChecker
}

func (q *questionChecker) checkWrong(sentence string) bool {

	result := q.checker.publishQueue(sentence)
	if result == 1 {
		return false
	} else {
		return true
	}
}

type answerChecker struct {
	checker *aiChecker
}

func newAnswerChecker() *answerChecker {

	newAnswerChecker := &answerChecker{}
	newAnswerChecker.checker = newAiChecker("answer", "ticker_check_mq_val.txt")
	return newAnswerChecker
}

func (a *answerChecker) checkWrong(sentence string) bool {

	result := a.checker.publishQueue(sentence)
	if result == 1 {
		return true
	} else {
		return false
	}
}
