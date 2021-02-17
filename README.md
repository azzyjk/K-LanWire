# K-LanWire
단비 AI 를 활용한 2021년도 건국대학교 해커톤
----
[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fazzyjk%2FK-LanWire&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)

# 프로젝트 개요
### 학교를 가지 못하는 학생들을 위한 대학생활 궁금증 해소 챗봇
#### 언택트 시대에 따라, 신입생을 포함한 많은 학생들이 대학생활 관련 정보를 얻을 수 없게 되었다. 이런 문제를 해결하고자, 대학생활 관련 정보룰 제공해주는 서비스의 필요성을 느끼게 되었다.
#### 이 서비스는 따라서, 단비 챗봇 AI를 이용해 선배가 알려주는 것처럼 대학셍활 전반 (요람, 학교생활, 수업 질문 등등)에 대한 정보를 제공해주는 것을 목표로 한다.


# 프로젝트 구조
![IMG-0384](https://user-images.githubusercontent.com/19836058/107953848-0956b680-6fdf-11eb-9093-5afe4c2f3b61.jpg)
### 프론트엔드
* React Native를 이용한 Cross-Platform (iOS, Android) 지원 어플리케이션
* User 정보를 저장하고, 데이터는 각각 단비 AI 서버와, 백엔드 서버에서 받아온다.
* 챗봇과 통신 / 질의응답 게시판 UI / 질의응답 게시판에서 대답 요청을 담당한다.

### Chatbot
* 단비 AI 플랫폼 기반
* 프론트엔드와 웹뷰 혹은 HTTP API 형식으로 연결되어있다.
* 대답이 없는 질문이 들어올 때, 백엔드에게 해당 질문을 넘겨준다.

### [백엔드](https://github.com/azzyjk/K-LanWire/tree/backend)
* Go 언어 기반으로 빠르고 부하분산에 유리하게 작성한다.
* 프론트엔드가 넘겨준 대답을 DB에 기록하고, 프론트엔드에게 질의응답 데이터를 넘겨준다.
* Chatbot에서 대답이 없는 질문을 DB에 기록하고, 복수의 대답 중 정제된 대답을 Chatbot 지식 데이터베이스로 갱신 요청한다.

### AI
  1. 악플 분류 AI 모델
      - 사용자가 작성한 답변이 부적절하면 필터링
      - KoELECTRA 모델을 Task에 맞게 Fine-Tune
      - 학습 데이터: https://github.com/kocohub/korean-hate-speech
      - KoELECTRA: https://github.com/monologg/KoELECTRA
  2. 질의 분류 AI 모델
      - 사용자가 질의응답 게시판에 질의한 내용이 실제 질문이 아닐 경우 필터링
      - KoELECTRA 모델을 Task에 맞게 Fine-Tune
      - 학습 데이터: https://github.com/warnikchow/3i4k

### DB
* MySQL로 되어있다.
* 질문과 대답 데이터셋을 저장 및 보관하고 있다.
