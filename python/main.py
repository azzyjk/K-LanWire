import os
import time
import pika
import hate_check as checker


_PATH = os.path.dirname(os.path.abspath(__file__)) + "/"


def GET_MQ_VALUE():
    return_val = {}
    with open(_PATH + "MQ_CRED.txt", "r", encoding='utf8') as mqfile:
        for line in mqfile:
            line = line.replace("\n", "").split("=")
            return_val[line[0]] = line[1]
    return return_val


class QueGetter:
    def __init__(self):
        get_mq_val = GET_MQ_VALUE()
        self._url = get_mq_val["MQ_URL"]
        self._port = get_mq_val["MQ_PORT"]
        self._vhost = get_mq_val["MQ_VHOST"]
        self._cred = pika.PlainCredentials(get_mq_val["MQ_ID"], get_mq_val["MQ_PW"])
        self._send_question_queue = get_mq_val["MQ_QUESTION_QUEUE"]
        self._send_answer_queue = get_mq_val["MQ_ANSWER_QUEUE"]
        self._recv_queue = get_mq_val["MQ_IN_QUEUE"]

    def receive_data(self):
        conn = pika.BlockingConnection(pika.ConnectionParameters(self._url, int(self._port), self._vhost, self._cred))
        channel = conn.channel()
        while True:
            raw_data = channel.basic_get(queue=self._recv_queue, auto_ack=True) # 3번째꺼가 원하고자 하는 String을 가져와서 처리를 시작한다.
            if raw_data[2] is not None:
                data = raw_data[2].decode().split("|")
                print("요청받은 pid : ", data[0], "  요청받은 항목 : ", data[1])
                checker_result = str(checker.predict([data[1]])[0])
                result = data[0].encode() + b'|' + checker_result.encode()
                # print("Final Result : ", result)
                if data[2] == "question":
                    channel.basic_publish(exchange='', routing_key=self._send_question_queue, body=result)
                if data[2] == "answer":
                    channel.basic_publish(exchange='', routing_key=self._send_answer_queue, body=result)

                # channel.basic_get(queue=self._recv_queue, auto_ack=True)  # 작업을 끝마친 후에서야 큐에서 작업을 지운다.
            time.sleep(0.1)


if __name__ == "__main__":
    test = QueGetter()
    test.receive_data()