package src

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
