package src

import (
	"database/sql"
	"encoding/json"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strings"
	"sync"
)

type serverHandler struct {
	db            *sql.DB
	showByMajor   map[string][3]string
	parsedMajor   []string
	rawMajor      string
	chatBotInput  *chatBotInputer
	answerChecker *answerChecker
}

func newServerHandler() *serverHandler {
	newHandler := &serverHandler{}
	newHandler.db, _ = sql.Open("mysql", "kuai:1q2w3e4r!@tcp(192.168.50.211:3306)/kudata")
	newHandler.showByMajor = make(map[string][3]string, 3)
	newHandler.showByMajor["컴퓨터공학부"] = [3]string{"컴퓨터공학부", "기계공학과", "전기전자공학부"}
	newHandler.showByMajor["기계공학과"] = [3]string{"기계공학과", "전기전자공학부", "컴퓨터공학부"}
	newHandler.showByMajor["전기전자공학부"] = [3]string{"전기전자공학부", "기계공학과", "기계공학과"}
	newHandler.parsedMajor = []string{"컴퓨터공학부", "기계공학과", "전기전자공학부"}
	newHandler.rawMajor = "cme"
	newHandler.chatBotInput = newChatBotInputer()
	newHandler.answerChecker = newAnswerChecker()

	return newHandler
}

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
	orderMajor := s.showByMajor[value[0]]
	queryData, err := s.db.Query("select Q.num, Q.questionEntranceYear, Q.questionDepartment, Q.question, Q.solved, A.answerEntranceYear, A.answerDepartment, A.answer, A.rank from question Q left join answer A on Q.num = A.questionNum order by Q.solved, FIELD(Q.questionDepartment, ?, ?, ?), Q.num, A.rank desc", orderMajor[0], orderMajor[1], orderMajor[2])
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

	// 1차 처리 - AI가 부적절한 답변이라고 생각하면 DB에 등록하지 않음
	if s.answerChecker.checkWrong(value[3]) {
		errorHandling(&serverResponse, 400, "부적절한 답변으로 간주되어 답변이 등록되지 않았습니다.", errors.New("답변 \""+value[3]+"\" 은 부적절합니다"))
		return
	}

	// 2차 처리 - answer, question 테이블에 값 넣기 (필수)
	var waiter sync.WaitGroup
	waiter.Add(2)

	go s.answerChatBotHandle(value)

	// answer 테이블에 값 넣기
	go func() {
		_, err := s.db.Exec("insert into answer (questionNum, answerEntranceYear, answerDepartment, answer, rank, answerId) values (?, ?, ?, ?, ?, ?)", value[0], value[1], value[2], value[3], value[4], "")
		if err != nil {
			errorHandling(&serverResponse, 500, "DB에 답변을 등록하는데 실패했습니다.", err)
			failed = true
		}
		waiter.Done()
	}()

	// uestion 테이블에 해당 값 solved로 표시
	go func() {
		_, err := s.db.Exec("update question set solved = ? where num = ?", 1, value[0])
		if err != nil {
			errorHandling(&serverResponse, 500, "DB에 질문 해결을 표시하는 데 실패했습니다.", err)
			failed = true
		}
		waiter.Done()
	}()

	// 3차 처리 - Final 응답
	waiter.Wait()
	if failed {
		return
	}
	serverResponse.WriteHeader(200)
	serverResponse.Write(messageInput("답변이 정상적으로 DB에 등록되었습니다."))
}

func (s *serverHandler) answerChatBotHandle(value []string) {
	// 답변을 확인하고 챗봇에 변형하거나 그대로 두기 위해 설정

	// 1차 처리 - DB에 질의해 답변의 존재여부 및 형태 알아옴
	status := 1
	var waiter sync.WaitGroup
	var isAnsweredValue [3]sql.NullString
	log.Println(value[0])
	isAnswered, err := s.db.Query("select num, rank, answerId from answer where questionNum = ? and answerId != ''", value[0])
	if !isAnswered.Next() { // 기존 답변이 등록되어있지 않을 때
		log.Println("기존 답변이 등록되어있지 않습니다.")
		status = 2
	} else {
		err = isAnswered.Scan(&isAnsweredValue[0], &isAnsweredValue[1], &isAnsweredValue[2])
		if err != nil {
			log.Println("DB의 기존 답변을 읽어오던 중 오류가 발생했습니다.")
		}
		if nullStringToString(isAnsweredValue[1]) == "0" && value[4] == "1" { // 기존에 등록된 답변이 학생이고 새로 등록하는 답변이 교수님일 때
			status = 3
		}
	}

	// 2차 처리 - 챗봇에 업데이트해야할 경우 업데이트, 새로 등록해야할 경우 새로 등록
	// 질문 및 대답으로 등록된 번호 가져오기
	if status != 1 { // 업데이트해야하는 경우에만 업데이트함
		waiter.Wait()
		waiter.Add(2)

		var question string
		var answerNum string

		go func() {
			reqVal, _ := s.db.Query("select question from question where num = ?", value[0])
			reqVal.Next()
			err = reqVal.Scan(&question)
			if err != nil {
				log.Println("질문 가져오는거 실패! : ", err)
			}
			waiter.Done()
		}()

		//key := []string{"questionNum", "answerEntranceYear", "answerDepartment", "answer", "rank"}
		go func() {
			reqVal, _ := s.db.Query("select num from answer where questionNum = ? and answerEntranceYear = ? and answerDepartment = ? and answer = ? and rank = ?", value[0], value[1], value[2], value[3], value[4])
			reqVal.Next()
			err = reqVal.Scan(&answerNum)
			if err != nil {
				log.Println("답변번호 가져오는거 실패! : ", err)
			}
			waiter.Done()
		}()

		waiter.Wait()

		// 여기서 챗봇 로직 업데이트
		//isAnswered, err := s.db.Query("select num, rank, answerId from answer where questionNum = ? and answerId != ''", value[0])

		if status == 2 {
			answerId := s.chatBotInput.addQAset(question, value[3])
			log.Println("정상등록됨 : ", answerId)
			_, err = s.db.Exec("update answer set answerId = ? where num = ?", answerId, answerNum)
			if err != nil {
				log.Println("DB에 챗봇 표시를 Marking하는 데에 실패했습니다.")
			}
		} else {
			waiter.Add(3)
			go func() {
				updated := s.chatBotInput.updateQAset(question, value[3], nullStringToString(isAnsweredValue[2]))
				if !updated {
					log.Println("대답 업데이트 실패!")
				}
				waiter.Done()
			}()
			go func() {
				_, err = s.db.Exec("update answer set answerId = '' where num = ?", nullStringToString(isAnsweredValue[0]))
				if err != nil {
					log.Println("기존 챗봇채택 대답 DB 업데이트 실패 : ", err)
				}
				waiter.Done()
			}()
			go func() {
				_, err2 := s.db.Exec("update answer set answerId = ? where num = ?", nullStringToString(isAnsweredValue[2]), answerNum)
				if err != nil {
					log.Println("새로운 챗봇채택 대답 DB 업데이트 실패 : ", err2)
				}
				waiter.Done()
			}()
			waiter.Wait()
		}
		log.Println("챗봇 및 DB 업데이트 종료")

	}
}
