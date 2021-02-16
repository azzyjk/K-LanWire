package src

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

type serverHandler struct {
	db            *sql.DB
	parsedMajor   []string
	rawMajor      string
	chatbotSender *http.Client
}

func newServerHandler() *serverHandler {
	newHandler := &serverHandler{}
	newHandler.db, _ = sql.Open("mysql", "kuai:1q2w3e4r!@tcp(192.168.50.211:3306)/kudata")
	newHandler.parsedMajor = []string{"컴퓨터공학부", "기계공학과", "전기전자공학부"}
	newHandler.rawMajor = "cme"
	newHandler.chatbotSender = &http.Client{}

	return newHandler
}

/*
// New QnAEnterHandler
func (s *serverHandler) qnAEnterHandler_NEW(serverResponse http.ResponseWriter, clientREquest *http.Request){

	// 1차 처리 - DB에서 값 긁어오기
	queryData, err := s.db.Query("select Q.num, Q.questionDepartment, Q.question, Q.solved, A.answerEntranceYear, A.answerDepartment, A.answer, A.rank from question Q left join answer A on Q.num = A.questionNum order by Q.num, A.num")
	if err != nil {
		errorHandling(&serverResponse, 500, "DB에서 값을 가져오는 도중 문제가 발생했습니다.", err)
		return
	}

	// 2차 처리 - DB 결과값 처리
	var qString [8]sql.NullString
	var qnaData qnaList

	for queryData.Next() {
		err = queryData.Scan(&qString[0], &qString[1], &qString[2], &qString[3], &qString[4], &qString[5], &qString[6], &qString[7], &qString[8])
		if err != nil {
			errorHandling(&serverResponse, 500, "DB에서 값을 파싱하던 도중 문제가 발생했습니다.", err)
			return
		}
		addQna(&qnaData, qString)
	}

	log.Println(value[0], qnaData)
	qnaData.Message = "정상 처리되었습니다"
	sendData, err := json.Marshal(qnaData)

	// 3차 처리 - Final 응답
	serverResponse.WriteHeader(200)
	serverResponse.Write(sendData)
}
*/

// 질의응답 게시판 입장 핸들링함
// 입력값 형태 : {"학과": "컴퓨터공학과"}
func (s *serverHandler) qnAEnterHandler(serverResponse http.ResponseWriter, clientRequest *http.Request) {

	if !(clientRequest.Method == "GET" || clientRequest.Method == "POST") {
		log.Println(clientRequest.Method)
		errorHandling(&serverResponse, 400, "GET만 받습니다.", errors.New("사용자요청이 GET이 아님"))
		return
	}

	// 요청 받는 부분 - 필요 없을 수 있음 (URL을 과마다 다르게 하면 필요없고, 동일하되 입력값에 따라서 다르게 처리할거면 필요함
	key := []string{"학과"}
	value := make([]string, len(key)) // value[0]만 학과를 의미함
	if readFromRequest(&serverResponse, clientRequest, key, &value) == false {
		return
	}

	// 1차 처리 - DB에서 값 긁어오기
	queryData, err := s.db.Query("select Q.num, Q.questionEntranceYear, Q.questionDepartment, Q.question, Q.solved, A.answerEntranceYear, A.answerDepartment, A.answer, A.rank from question Q left join answer A on Q.num = A.questionNum where Q.questionDepartment = '" + value[0] + "' order by Q.num, A.num")
	if err != nil {
		errorHandling(&serverResponse, 500, "DB에서 값을 가져오는 도중 문제가 발생했습니다.", err)
		return
	}

	// 2차 처리 - DB 결과값 처리
	var qString [9]sql.NullString
	var qnaData qnaList

	for queryData.Next() {
		err = queryData.Scan(&qString[0], &qString[1], &qString[2], &qString[3], &qString[4], &qString[5], &qString[6], &qString[7], &qString[8])
		if err != nil {
			errorHandling(&serverResponse, 500, "DB에서 값을 파싱하던 도중 문제가 발생했습니다.", err)
			return
		}
		addQna(&qnaData, qString)
	}

	log.Println(value[0], qnaData)
	qnaData.Message = "정상 처리되었습니다"
	sendData, err := json.Marshal(qnaData)

	// 3차 처리 - Final 응답
	serverResponse.WriteHeader(200)
	serverResponse.Write(sendData)
}

// 챗봇에서 질문을 등록할 때 핸들링함
// 입력값 형태 : {"question": "전공이수가 몇학점이죠?", "id": "c17"}
func (s *serverHandler) questionRegisterHandler(serverResponse http.ResponseWriter, clientRequest *http.Request) {

	if clientRequest.Method != "POST" {
		errorHandling(&serverResponse, 400, "POST만 받습니다.", errors.New("사용자요청이 POST이 아님"))
		return
	}

	// 요청 받아서 값 읽어오는 부분
	key := []string{"question", "id"}
	value := make([]string, len(key)) // value[0]만 학과를 의미함
	if readFromRequest(&serverResponse, clientRequest, key, &value) == false {
		return
	}

	// 1차 처리 - 읽어온 값에서 DB에 넣을 값으로 변환
	log.Println("index : ", strings.Index(s.rawMajor, string(value[1][0])))
	log.Println(s.parsedMajor[1])
	major := s.parsedMajor[strings.Index(s.rawMajor, string(value[1][0]))]
	year := "20" + value[1][1:]

	// 2차 처리 - DB에 해당 값을 넣음
	_, err := s.db.Exec("insert into question(questionEntranceYear, questionDepartment, question) values (?, ?, ?)", year, major, value[0])
	if err != nil {
		errorHandling(&serverResponse, 500, "DB에 값을 insert하는 데 오류가 발생했습니다", err)
		return
	}

	// 3차 처리 - Final 응답
	serverResponse.WriteHeader(200)
	serverResponse.Write(messageInput("질문이 정상적으로 DB에 등록되었습니다."))
}

// 질의응답 게시판에서 질문에 응답할 때 핸들링함
// 입력값 형태 : {"questionNum": "1", "answerEntranceYear": "2019", "answerDepartment": "컴퓨터공학부", "answer": "20학점입니다", "rank": "0"}
func (s *serverHandler) answerRegisterHandler(serverResponse http.ResponseWriter, clientRequest *http.Request) {

	if clientRequest.Method != "POST" {
		errorHandling(&serverResponse, 400, "POST만 받습니다.", errors.New("사용자요청이 POST이 아님"))
		return
	}

	// 요청 받는 부분 -> 값까지 정상적으로 읽어옴
	key := []string{"questionNum", "answerEntranceYear", "answerDepartment", "answer", "rank"}
	value := make([]string, len(key)) // value[0]만 학과를 의미함
	if readFromRequest(&serverResponse, clientRequest, key, &value) == false {
		return
	}
	failed := false

	var waiter sync.WaitGroup
	waiter.Add(3)

	// 1차 처리 - answer 테이블에 값 넣기
	go func() {
		_, err := s.db.Exec("insert into answer (questionNum, answerEntranceYear, answerDepartment, answer, rank) values (?, ?, ?, ?, ?)", value[0], value[1], value[2], value[3], value[4])
		if err != nil {
			errorHandling(&serverResponse, 500, "DB에 답변을 등록하는데 실패했습니다.", err)
			failed = true
		}
		waiter.Done()
	}()

	// 2차 처리 - question 테이블에 해당 값 solved로 표시
	go func() {
		_, err := s.db.Exec("update question set solved = ? where num = ?", 1, value[0])
		if err != nil {
			errorHandling(&serverResponse, 500, "DB에 질문 해결을 표시하는 데 실패했습니다.", err)
			failed = true
		}
		waiter.Done()
	}()

	// 3차 처리 - Chatbot에 답변 등록하기
	go func() {
		var question string
		reqVal, _ := s.db.Query("select question from question where num = ?", value[0])
		reqVal.Next()
		err := reqVal.Scan(&question)
		if err != nil {
			log.Println("질문 가져오는거 실패! : ", err)
		}
		go s.sendQAsetToChatbot(question, value[3])
		waiter.Done()
	}()

	// 4차 처리 - Final 응답
	waiter.Wait()
	if failed {
		return
	}
	serverResponse.WriteHeader(200)
	serverResponse.Write(messageInput("답변이 정상적으로 DB에 등록되었습니다."))
}

func (s *serverHandler) sendQAsetToChatbot(question, answer string) {

	value := "{\"q\": \"" + question + "\", \"a\": \"" + answer + "\"}"
	byteValue := []byte(value)

	request, err := http.NewRequest("POST", "https://danbee-21konkuk.mybluemix.net/qa", bytes.NewBuffer(byteValue))
	if err != nil {
		panic("Failed to create Request")
	}

	request.Header.Add("Content-type", "application/json")

	response, err := s.chatbotSender.Do(request)
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
