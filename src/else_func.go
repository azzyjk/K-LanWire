package src

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func readFromRequest(response *http.ResponseWriter, request *http.Request, key []string, reqValue *[]string) bool {

	responseResult, err := ioutil.ReadAll(request.Body) // 1. url에 받아온 값을 raw string으로 Parsing함.
	log.Println("Raw INPUT : ", string(responseResult))
	if string(responseResult) == "" {
		errorHandling(response, 400, "빈 데이터를 요청했습니다.", errors.New("빈 데이터가 들어옴"))
		return false
	}
	var test map[string]interface{}
	_ = json.Unmarshal(responseResult, &test)
	for index, data := range key {
		(*reqValue)[index] = test[data].(string)
		if (*reqValue)[index] == "" {
			errorHandling(response, 402, data+"를 받아오는 데 실패했습니다.", err)
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

	log.Println("error occured! : ", err, "\t", errMsg)
	(*serverResponse).WriteHeader(code)
	(*serverResponse).Write([]byte("{\"message\": \"" + errMsg + "\"}"))
}

func messageInput(message string) []byte {
	return []byte("{\"message\": \"" + message + "\"}")
}
