VERSION=v1.0.3
BUILD_DATE=2024-02-29

build:
	go build -o ./examples/bin/main ./main.go

test:
	cd ./examples && ./examples/bin/main init

install:
	go install -ldflags "-X 'main.Version=${v1.0.3}' -X 'main.BuildDate=${BUILD_DATE}'"

push:
	git push
	git tag ${VERSION}
	git push origin --tags
