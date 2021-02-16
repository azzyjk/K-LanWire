package src

import (
	"bytes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
	"log"
	"strconv"
	"time"
)

type aiChecker struct {

	// 구조체 고유 이름
	name string

	// rabbitMQ 설정 시 필요한 값
	url       string
	port      string
	vhost     string
	cred      string
	sendQueue string
	recvQueue string

	//recvQueue connection과 channel
	recvQueueConnection *amqp.Connection
	recvQueueChannel    *amqp.Channel

	//sendQueue connection과 channel
	sendQueueConnection *amqp.Connection
	sendQueueChannel    *amqp.Channel

	//
	connector             [3]chan int
	connectorEmptyChecker chan int
	differ                int
}

func newAiChecker(name, credFileName string) *aiChecker {
	newAiChecker := &aiChecker{}
	//newAiChecker.db, _ = sql.Open("mysql", "kuai:1q2w3e4r!@tcp(192.168.50.211:3306)/kudata")
	newAiChecker.name = name

	queueCredInfo := readFromFile("/" + credFileName)
	newAiChecker.url = queueCredInfo["MQ_URL"]
	newAiChecker.port = queueCredInfo["MQ_PORT"]
	newAiChecker.vhost = queueCredInfo["MQ_VHOST"]
	newAiChecker.cred = queueCredInfo["MQ_ID"] + ":" + queueCredInfo["MQ_PW"]
	newAiChecker.sendQueue = queueCredInfo["MQ_OUT_QUEUE"]
	newAiChecker.recvQueue = queueCredInfo["MQ_IN_QUEUE"]

	newAiChecker.recvQueueConnection, newAiChecker.recvQueueChannel = newAiChecker.getConnection()
	newAiChecker.sendQueueConnection, newAiChecker.sendQueueChannel = newAiChecker.getConnection()

	//sender와 consumer를 잇는 connector 설정
	newAiChecker.connectorEmptyChecker = make(chan int, 3)
	for i := 0; i < 3; i++ {
		newAiChecker.connector[i] = make(chan int, 1)
		i := i
		go func() { newAiChecker.connectorEmptyChecker <- i }()
	}
	rawVal := strconv.Itoa(int(time.Now().Unix()))[5:]
	val, _ := strconv.Atoi(rawVal)
	newAiChecker.differ = val * 10

	//goroutine으로 미리 돌음
	go newAiChecker.consumeQueue()

	return newAiChecker
}

func (a *aiChecker) getConnection() (*amqp.Connection, *amqp.Channel) {
	//Timelog("amqp://" + q.cred + "@" + q.url + ":" + q.port + "/" + q.vhost)
	connection, err := amqp.Dial("amqp://" + a.cred + "@" + a.url + ":" + a.port + "/" + a.vhost)
	if err != nil {
		panic("connection 연결 오류")
	}

	channel, err := connection.Channel()
	if err != nil {
		panic("channel 얻기 오류")
	}

	return connection, channel
}

func (a *aiChecker) consumeQueue() {

	//Timelog(queueName)
	val1, _ := a.recvQueueChannel.Consume(a.recvQueue, "", true, true, true, false, nil)
	amqpHandler := AMQPHandler{queDelivery: val1}

	for {
		receiveVal := bytes.Split(amqpHandler.GetFromQueue(), []byte("|"))
		// 문장에 대한 대답이므로 1이면 부정, 0이면 긍정인 단순한 int 형으로 return
		// 0번은 각자 프로세스이니, (전달할 때 고유 ID)|(긍부정) 식으로 한다고 가정
		roomNum, err := strconv.Atoi(string(receiveVal[0]))
		roomNum = roomNum - a.differ
		if err != nil {
			panic("숫자 변환 오류")
		}
		result, _ := strconv.Atoi(string(receiveVal[1]))

		if roomNum > 3 {
			log.Println("현재 프로그램에서 Publish한 값이 아니므로 무시합니다.")
		} else {
			a.connector[roomNum] <- result
		}
	}
}

func (a *aiChecker) publishQueue(value string) int {

	for {
		roomNum := <-a.connectorEmptyChecker
		passValue := strconv.Itoa(roomNum+a.differ) + "|" + value + "|" + a.name

		err := a.sendQueueChannel.Publish(
			"",
			a.sendQueue,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(passValue),
			},
		)
		if err != nil {
			panic("Queue publishing 과정 중 문제 발생")
		}

		queueData := 2
		isBreak := false
		timer := time.NewTimer(time.Second * 20)
		for {
			select {
			case data := <-a.connector[roomNum]:
				queueData = data
				isBreak = true
			case <-timer.C:
				isBreak = true
			default:
			}
			if isBreak {
				timer.Stop()
				break
			}
		}

		// 데이터를 받아온 뒤에야 빈 방이 있다고 다시 알려줌
		a.connectorEmptyChecker <- roomNum

		// 이후 데이터가 잘못되면 재시도, 아니면 정상 값을 return
		// 일단 그냥 값 return
		return queueData
	}

	return 0
}
