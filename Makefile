IMG     := mramshaw4docs/restfulapi
VERSION := 1.0.0
GOPATH  := "`pwd`/go"
GOOS    := linux
GOARCH  := amd64

all:	deps build

env:
	echo $(GOPATH)

deps:
	mkdir -p $(GOPATH)
	GOPATH=$(GOPATH) go get -d -v .

build:
	GOPATH=$(GOPATH) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -v -o RestfulGorillaMux .

package:
	docker build -t $(IMG) .
	docker tag $(IMG) $(IMG):$(VERSION)

push:
	docker push $(IMG)

clean:
	rm -rf $(GOPATH)
	rm RestfulGorillaMux
