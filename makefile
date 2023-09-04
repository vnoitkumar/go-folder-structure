format:
	go fmt ./...

test:
	go test ./... -short -count=1 -p=1

start:
	go run main.go

build:
	go build

health_check:
	open -a "Google chrome" http://localhost:${PORT}/health-check

run:
	CompileDaemon -command="./learning-go-lang"
