IMG       := mramshaw4docs/restfulapi
VERSION   := 1.0.2
GOPATH    := "`pwd`/go"
GOOS      := linux
GOARCH    := amd64

PORT      := 8100

# swagger-ui
CORS_HOST := http://localhost:3200

all:	deps test build run

env:
	echo $(GOPATH)

deps:
	mkdir -p $(GOPATH)
	GOPATH=$(GOPATH) go get -d -v .

test:
	GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) go test -v ./api/people

build:
	GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o RestfulGorillaMux .

run:
	GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) PORT=$(PORT) CORS_HOST=$(CORS_HOST) ./RestfulGorillaMux

package:
	docker build -t $(IMG) .
	docker tag $(IMG) $(IMG):$(VERSION)

push:
	docker push $(IMG)

clean:
	rm -rf $(GOPATH)
	rm RestfulGorillaMux
