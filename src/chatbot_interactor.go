package src

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type chatBotInputer struct {
	chatbotSender *http.Client

	addQAsetUrl    string
	updateQAsetUrl string
	deleteQAsetUrl string
}

func newChatBotInputer() *chatBotInputer {
	newChatBotInputer := &chatBotInputer{}
	newChatBotInputer.chatbotSender = &http.Client{}
	newChatBotInputer.addQAsetUrl = "https://danbee-21konkuk.mybluemix.net/qa"
	newChatBotInputer.updateQAsetUrl = "https://danbee-21konkuk.mybluemix.net/qa_update"
	newChatBotInputer.deleteQAsetUrl = "https://danbee-21konkuk.mybluemix.net/qa_del"
	return newChatBotInputer
}

func (c *chatBotInputer) basicRequest(key, value []string, reqUrl string) map[string]interface{} {

	var parsedValue map[string]interface{}

	sendVal := "{"
	for index := range key {
		sendVal += "\"" + key[index] + "\": " + value[index] + ","
	}
	sendVal = sendVal[:len(sendVal)-1]
	sendVal += "}"
	byteSendVal := []byte(sendVal)

	request, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(byteSendVal))
	if err != nil {
		log.Println("Failed to create Request")
	}
	request.Header.Add("Content-type", "application/json")

	response, err := c.chatbotSender.Do(request)
	if err != nil {
		log.Println("Failed to receive Data, check server status")
	}

	byteResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Failed to parse data")
	}

	err = json.Unmarshal(byteResponse, &parsedValue)
	if err != nil {
		log.Println("Failed to parse data - convert to json : ", err)
	}

	return parsedValue
}

func (c *chatBotInputer) addQAset(question, answer string) string {

	key := []string{"q", "a"}
	value := []string{"\"" + question + "\"", "\"" + answer + "\""}
	result := c.basicRequest(key, value, c.addQAsetUrl)
	return strconv.Itoa(int(result["id"].(float64)))
}

func (c *chatBotInputer) updateQAset(question, answer, id string) bool {

	key := []string{"q", "a", "id"}
	value := []string{"\"" + question + "\"", "\"" + answer + "\"", id}
	result := c.basicRequest(key, value, c.updateQAsetUrl)
	if result["update"].(float64) == -1 {
		return false
	} else {
		return true
	}
}

func (c *chatBotInputer) deleteQAset(id string) bool {

	key := []string{"id"}
	value := []string{id}
	result := c.basicRequest(key, value, c.deleteQAsetUrl)
	if result["del"].(float64) == -1 {
		return false
	} else {
		return true
	}
}
