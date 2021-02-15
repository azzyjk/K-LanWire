package src

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type AiChecker struct {
	db *sql.DB
}

func newAiChecker() *AiChecker {
	aiChecker := &AiChecker{}
	aiChecker.db, _ = sql.Open("mysql", "kuai:1q2w3e4r!@tcp(192.168.50.211:3306)/kudata")

	return aiChecker
}

func (a *AiChecker) aiCheckAnswer() {

	ticker := time.NewTicker(time.Hour)

	for {
		<-ticker.C
		go a.oneCheck()
	}
}

func (a *AiChecker) oneCheck() {

	log.Println("한번 호출됨")
}
