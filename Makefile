VERSION=v1.0.7

build:
	go build -o ./examples/bin/main ./main.go

test:
	cd ./examples && ./examples/bin/main init

install:
	go build
	go install

test:
	go run main.go migrate --host localhost:3000 --path /database/m 

push:
	git push
	git tag ${VERSION}
	git push origin --tags
