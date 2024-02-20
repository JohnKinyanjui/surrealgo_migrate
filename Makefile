build:
	go build -o ./examples/bin/main ./main.go

test:
	cd ./examples && ./examples/bin/main init

install:
	go install