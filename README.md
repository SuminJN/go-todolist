# TO DO List
Go 언어로 작성된 CLI 기반 할일 관리 프로그램입니다.

## 기능 요구사항

- [x] 할일을 추가할 수 있다
- [x] 할일 목록을 조회할 수 있다
- [x] 할일을 완료 처리할 수 있다
- [x] 할일을 삭제할 수 있다
- [x] 데이터를 JSON 파일에 영구 저장한다

## 기능 상세

### 1. 할일 추가
- 사용자가 입력한 제목으로 새로운 할일을 생성한다
- 각 할일에는 고유한 ID가 자동으로 부여된다
- 새로 추가된 할일은 미완료 상태로 저장된다

### 2. 할일 목록 조회
- 저장된 모든 할일을 출력한다
- 완료된 할일과 미완료 할일을 구분하여 표시한다
- 할일이 없는 경우 적절한 메시지를 출력한다

### 3. 할일 완료 처리
- ID를 입력받아 해당 할일을 완료 상태로 변경한다
- 존재하지 않는 ID인 경우 오류 메시지를 출력한다

### 4. 할일 삭제
- ID를 입력받아 해당 할일을 목록에서 제거한다
- 존재하지 않는 ID인 경우 오류 메시지를 출력한다

### 5. 데이터 저장
- 모든 할일은 `todos.json` 파일에 저장된다
- 프로그램 재실행 시 이전 데이터를 불러온다
- 파일이 없는 경우 새로 생성한다

## 동작 방식

1. 사용자가 명령어를 입력
2. 명령어를 파싱하고 적절한 함수 호출
3. JSON 파일에서 데이터 읽기/쓰기
4. commands.go의 각 함수가 할일 추가, 조회, 완료, 삭제 처리
5. 구조체로 할일 데이터를 표현

## Go 학습 포인트

- 함수 정의
- 구조체
- 슬라이스 문법
- 에러 처리
- 포인터
- JSON 처리
- 문자열

## 실행 화면

### 사용법 출력

<img width="1188" height="340" alt="image" src="https://github.com/user-attachments/assets/d86628cf-dd40-4f5c-a666-ebd4225e638e" />


### 새로운 할일 추가

<img width="1502" height="104" alt="image" src="https://github.com/user-attachments/assets/e6642961-4f4b-49d6-a012-b5b1cf0f65ce" />


### 할일 목록 보기

<img width="1272" height="382" alt="image" src="https://github.com/user-attachments/assets/91e85105-b326-42ae-97a3-7bba3c6f8326" />


### 할일 완료 표시

<img width="1316" height="550" alt="image" src="https://github.com/user-attachments/assets/27de63da-77a1-42d9-ae6a-32efcf3012ba" />



### 할일 삭제

<img width="1344" height="434" alt="image" src="https://github.com/user-attachments/assets/ef847ce5-5c02-4b2c-8988-0177f325e4d1" />


### 저장된 JSON 데이터

<img width="748" height="1006" alt="image" src="https://github.com/user-attachments/assets/6bef2f14-170e-45e0-b3d3-a9e5da309409" />

