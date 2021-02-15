package src

import (
	"encoding/json"
	"log"
	"net/http"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type serverHandler struct {
	db *sql.DB
}

func newServerHandler() *serverHandler {
	newHandler := &serverHandler{}

	newHandler.db, _ = sql.Open("mysql", "kuai:1q2w3e4r!@tcp(192.168.50.211:3306)/kudata")

	return newHandler
}

// 로그인 부분 핸들러 -> 필요 없음
func (s *serverHandler) loginHandler(serverResponse http.ResponseWriter, clientRequest *http.Request) {

	// 요청 받는 부분
	requestBody, err := readFromRequest(clientRequest)
	if err != nil {
		log.Println("error occured : ", err)
		serverResponse.WriteHeader(400)
		return
	}
	log.Println(requestBody)

	// 1차 처리 - 값의 형식이 적절한지 확인

	// 2차 처리 - DB에 사용자 정보 입력

	// 3차 처리 - DB의 응답을 파악

	// 4차 처리 - Final 응답
	serverResponse.WriteHeader(200) // -> Parsing에 따른 값의 다른 전달
	serverResponse.Write([]byte("{\"response\": \"나는 밥 잘 먹고 있어!\"}"))
}

// 챗봇 API 입력 핸들러 -> 필요 없음
func (s *serverHandler) chatbotRequestHandler(serverResponse http.ResponseWriter, clientRequest *http.Request) {

	// 요청 받는 부분
	requestBody, err := readFromRequest(clientRequest)
	if err != nil {
		log.Println("error occured : ", err)
		serverResponse.WriteHeader(400)
		return
	}
	log.Println(requestBody)

	// 1차 처리 - 값의 형식이 적절한지 확인

	// 2차 처리 - 챗봇에 질의

	// 3차 처리 - 챗봇의 응답을 파악

	// 4차 처리 - Final 응답
	serverResponse.WriteHeader(200)
	serverResponse.Write([]byte("{\"response\": \"나는 밥 잘 먹고 있어!\"}"))
}

// 질의응답 게시판 입장 핸들링함 -> 입력값으로 학과 받아야함
func (s *serverHandler) qnAEnterHandler(serverResponse http.ResponseWriter, clientRequest *http.Request) {

	// 요청 받는 부분 - 필요 없을 수 있음 (URL을 과마다 다르게 하면 필요없고, 동일하되 입력값에 따라서 다르게 처리할거면 필요함
	requestBody, err := readFromRequest(clientRequest)
	if err != nil {
		log.Println("error occured : ", err)
		serverResponse.WriteHeader(400)
		return
	}
	log.Println("reqbody : ", requestBody)

	// 1차 처리 - DB에서 값 긁어오기
	queryData, err := s.db.Query("select Q.num, Q.questionEntranceYear, Q.questionDepartment, Q.question, Q.solved, A.answerEntranceYear, A.answerDepartment, A.answer from question Q left join answer A on Q.num = A.questionNum where Q.questionDepartment = '" + requestBody + "' order by Q.num, A.num")
	if err != nil {
		log.Println("error occured : ", err)
		serverResponse.WriteHeader(500)
		serverResponse.Write([]byte("{\"message\": \"DB에서 값을 가져오는 도중 문제가 발생했습니다. (과 입력 문제 가능성 있음)\"}"))
		return
	}

	// 2차 처리 - DB 결과값 처리
	var qString [8]sql.NullString
	var qnaData qnaList

	for queryData.Next() {
		err = queryData.Scan(&qString[0], &qString[1], &qString[2], &qString[3], &qString[4], &qString[5], &qString[6], &qString[7])
		if err != nil {
			log.Println("error occured : ", err)
			serverResponse.WriteHeader(500)
			serverResponse.Write([]byte("{\"message\": \"DB에서 값을 파싱하던 도중 문제가 발생했습니다.\"}"))
			return
		}
		log.Println(qString[0], qString[1], qString[2], qString[3], qString[4], qString[5], qString[6], qString[7])
		addQna(&qnaData, qString)
	}

	log.Println(qnaData)
	qnaData.Message = "정상 처리되었습니다"
	sendData, err := json.Marshal(qnaData)
	//log.Println(sendData)

	// 3차 처리 - Final 응답
	serverResponse.WriteHeader(200)
	serverResponse.Write(sendData)
}

// 질의응답 게시판에서 질문에 응답할 때 핸들링함
func (s *serverHandler) answerRegisterHandler(serverResponse http.ResponseWriter, clientRequest *http.Request) {

	// 요청 받는 부분
	requestBody, err := readFromRequest(clientRequest)
	if err != nil {
		log.Println("error occured : ", err)
		serverResponse.WriteHeader(400)
		return
	}
	log.Println(requestBody)

	// 1차 처리 - 값의 형식이 적절한지 확인

	// 2차 처리 - 값 전처리 (AI가 들어갈수도 있고, 아닐수도 있음)

	// 3차 처리 - 챗봇에 해당 값을 입력
	// -> goroutine으로 별도로 처리해서, 챗봇 단계에서 오류가 나더라도 별도로 처리할 수 있도록 함

	// 4차 처리 - Final 응답
	serverResponse.WriteHeader(200)
	serverResponse.Write([]byte("{\"response\": \"나는 밥 잘 먹고 있어!\"}"))
}

// 챗봇에서 질문을 등록할 때 핸들링함
func (s *serverHandler) questionRegisterHandler(serverResponse http.ResponseWriter, clientRequest *http.Request) {

	// 요청 받는 부분
	requestBody, err := readFromRequest(clientRequest)
	if err != nil {
		log.Println("error occured : ", err)
		serverResponse.WriteHeader(400)
		return
	}
	log.Println(requestBody)

	// 1차 처리 - 값의 형식이 적절한지 확인

	// 2차 처리 - 값 전처리 (AI가 들어갈수도 있고, 아닐수도 있음)

	// 3차 처리 - DB에 해당 값을 저장

	// 4차 처리 - Final 응답
	serverResponse.WriteHeader(200)
	serverResponse.Write([]byte("{\"response\": \"나는 밥 잘 먹고 있어!\"}"))
}
