Backend RESTful API 사용방법
----

#### 1. 질의응답 게시판 데이터 접근 
   * url : http://119.192.41.219:9184/qna_enter
   * method : GET
   * parameter : {"학과": "컴퓨터공학부"}

    * 코드예제
```python
showQnA = {"학과": "컴퓨터공학부"}
req = requests.get("http://119.192.41.219:9184/qna_enter", data=showQnA)
print(req.text)
```
***

#### 2. 질문 등록
   * url : http://119.192.41.219:9184/qna_question
   * method : POST
   * parameter : {"question": "테스트 중입니다.", "questionEntranceYear": "2020", "questionDepartment": "컴퓨터공학부"}
     
    * 코드예제
```python
addQ = {"question": "테스트 중입니다.", "questionEntranceYear": "2020", "questionDepartment": "컴퓨터공학부"}
req = requests.post("http://119.192.41.219:9184/qna_question", data=addQ)
print(req.text)
```
***

#### 3. 답변 등록
   * url : http://119.192.41.219:9184/qna_answer
   * method : POST
   * parameter : {"questionNum": "9", "answerEntranceYear": "2019", "answerDepartment": "컴퓨터공학부", "answer": "테스트가 성공적으로 표시됩니다.", "rank": "1"}
    *코드예제
```python     
addA = {"questionNum": "9", "answerEntranceYear": "2019", "answerDepartment": "컴퓨터공학부", "answer": "테스트가 성공적으로 표시됩니다.", "rank": "1"}
req = requests.post("http://119.192.41.219:9184/qna_answer", data=addA)
print(req.text)
```