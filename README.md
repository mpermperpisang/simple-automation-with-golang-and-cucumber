### Installation
- [x] Install Go - https://golang.org/dl/
- [x] Install Godog - https://github.com/cucumber/godog
- [x] Install env - https://github.com/joho/godotenv
or
- [x] go mod download

### Env
- [x] `cp env.sample .env`

### Run
- [x] `godog --tags=@example --format=cucumber > test/report/cucumber_report.json`</br>
or
- [x] `godog -t=@example -f=cucumber > test/report/cucumber_report.json`
