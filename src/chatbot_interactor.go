package src

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

type chatBotInputer struct {
	chatbotSender *http.Client
}

func newChatBotInputer() *chatBotInputer {
	newChatBotInputer := &chatBotInputer{}
	newChatBotInputer.chatbotSender = &http.Client{}
	return newChatBotInputer
}

func (c *chatBotInputer) sendQAsetToChatbot(question, answer string) {

	value := "{\"q\": \"" + question + "\", \"a\": \"" + answer + "\"}"
	byteValue := []byte(value)

	request, err := http.NewRequest("POST", "https://danbee-21konkuk.mybluemix.net/qa", bytes.NewBuffer(byteValue))
	if err != nil {
		panic("Failed to create Request")
	}

	request.Header.Add("Content-type", "application/json")

	response, err := c.chatbotSender.Do(request)
	if err != nil {
		panic("Failed to receive Data, check server status")
	}

	byteResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic("Failed to received data")
	}

	if string(byteResponse) == question {
		log.Println("질문 : ", question, "\t대답 : ", answer, " 이 정상적으로 등록되었습니다.")
	} else {
		log.Println("정상적으로 등록되지 않았습니다.")
	}
}
