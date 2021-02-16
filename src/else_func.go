package src

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
)

func readFromRequest(response *http.ResponseWriter, request *http.Request, key []string, reqValue *[]string) bool {

	responseResult, err := ioutil.ReadAll(request.Body) // 1. url에 받아온 값을 raw string으로 Parsing함.
	log.Println("URL :", request.URL.EscapedPath(), "\tIP : ", request.RemoteAddr, "\tINPUT : ", string(responseResult))
	if string(responseResult) == "" {
		errorHandling(response, 400, "빈 데이터를 요청했습니다.", errors.New("빈 데이터가 들어옴"))
		return false
	}
	var test map[string]interface{}
	_ = json.Unmarshal(responseResult, &test)
	for index, data := range key {
		if reflect.TypeOf(test[data]) != reflect.TypeOf("str") {
			errorHandling(response, 400, "값이 형식에 맞지 않습니다.", errors.New("JSON 값이 요청한 형식에 맞지 않음"))
			return false
		}
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

func readFromFile(fileName string) map[string]string {

	curLoc, _ := os.Getwd()
	dbData, err := ioutil.ReadFile(curLoc + fileName)
	if err != nil {
		log.Println("파일 불러오기 실패 : ", err)
	}

	var fileMap map[string]string
	fileMap = make(map[string]string)
	lines := strings.Split(string(dbData), "\n")
	//Timelog(lines)
	for i := 0; i < len(lines)-1; i++ {
		singleLine := strings.Split(lines[i], "=")
		fileMap[singleLine[0]] = singleLine[1]
	}

	return fileMap
}
