package src

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func readFromRequest(response *http.ResponseWriter, request *http.Request, key []string, reqValue *[]string) bool {

	responseResult, err := ioutil.ReadAll(request.Body) // 1. url에 받아온 값을 raw string으로 Parsing함.
	if err != nil {
		errorHandling(response, 400, "요청값을 Parsing하는 데 실패했습니다. {\"학과\": \"컴퓨터공학부\"} 와 같은 형식으로 보내주세요", err)
		return false
	}
	keyValueSet, err := url.ParseQuery(string(responseResult)) // 2. string을 url.ParseQuery를 이용해 원래 요청값으로 다시 Parsing
	if err != nil {
		errorHandling(response, 400, "전달된 body를 해석하지 못했습니다.", err)
		return false
	}
	for index, data := range key {
		(*reqValue)[index] = keyValueSet.Get(data)
		if (*reqValue)[index] == "" {
			errorHandling(response, 400, data+"를 받아오는 데 실패했습니다.", err)
			return false
		}
	}
	return true
}

func nullStringToString(inputString sql.NullString) string {

	if inputString.Valid {
		return inputString.String
	} else {
		return ""
	}
}

func errorHandling(serverResponse *http.ResponseWriter, code int, errMsg string, err error) {

	log.Println(err)
	(*serverResponse).WriteHeader(code)
	(*serverResponse).Write([]byte("{\"message\": \"" + errMsg + "\"}"))
}

func messageInput(message string) []byte {
	return []byte("{\"message\": \"" + message + "\"}")
}
