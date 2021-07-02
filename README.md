### Installation
- [x] Install Go - https://golang.org/dl/
- [x] Add export .bash_profile (or any profile you usually use)
```
export GOROOT=/usr/local/go
export GOPATH=$HOME/Documents/project/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
- [x] Install Godog - https://github.com/cucumber/godog
- [x] Install env - https://github.com/joho/godotenv</br>
or
- [x] go mod download

### Env
- [x] `cp env.sample .env`

### Run
- [x] `godog --tags=@example --format=cucumber > test/report/cucumber_report.json`</br>
or
- [x] `godog -t=@example -f=cucumber > test/report/cucumber_report.json`
