package src

import (
	"log"
	"net/http"
)

func ServerInit() {

	// set flag
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	handler := newServerHandler()
	http.HandleFunc("/qna_enter", handler.qnAEnterHandler)
	http.HandleFunc("/qna_answer", handler.answerRegisterHandler)
	http.HandleFunc("/qna_question", handler.questionRegisterHandler)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Println("ListenAndServe", err)
	}
}

func AiCheckerInit() {

	// set flag
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	checker := newAiChecker()
	checker.aiCheckAnswer()
}
