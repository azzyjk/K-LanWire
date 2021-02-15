package src

import (
	"database/sql"
	"strconv"
)

type qnaList struct {
	Message      string
	QuestionList []qaDataset
}

type qaDataset struct {
	Num          int64
	QuestionYear int64
	Question     string
	Solved       bool
	Answers      []answer
}

type answer struct {
	Answer     string
	AnswerYear int64
}

func addQna(qnaData *qnaList, queryData [8]sql.NullString) {

	// 만약 처음으로 질문을 넣을 떄
	if len(qnaData.QuestionList) == 0 {
		// 새로운 질문 넣기
		if addOnlyQuestion(qnaData, queryData) == true {
			return
		}

		// 대답
		addOnlyAnswer(qnaData, queryData)
		return
	}

	// 들어온 새로운 것이 대답뿐일 떄
	if strconv.FormatInt(qnaData.QuestionList[len(qnaData.QuestionList)-1].Num, 10) == nullStringToString(queryData[0]) {
		addOnlyAnswer(qnaData, queryData)
	} else {
		if addOnlyQuestion(qnaData, queryData) == true {
			return
		}
		addOnlyAnswer(qnaData, queryData)
	}
}

func addOnlyAnswer(qnaData *qnaList, queryData [8]sql.NullString) {
	var ans answer
	ans.Answer = nullStringToString(queryData[7])
	ans.AnswerYear, _ = strconv.ParseInt(nullStringToString(queryData[5]), 10, 64)
	qnaData.QuestionList[0].Answers = append(qnaData.QuestionList[0].Answers, ans)
}

// -> 질문만 추가하는게 참이면, return하면 됨
func addOnlyQuestion(qnaData *qnaList, queryData [8]sql.NullString) bool {
	var qna qaDataset
	qna.Num, _ = strconv.ParseInt(nullStringToString(queryData[0]), 10, 64)
	qna.QuestionYear, _ = strconv.ParseInt(nullStringToString(queryData[1]), 10, 64)
	qna.Question = nullStringToString(queryData[3])
	if nullStringToString(queryData[4]) == "1" {
		qna.Solved = true
	} else {
		qna.Solved = false
	}
	qnaData.QuestionList = append(qnaData.QuestionList, qna)
	return !qna.Solved
}
