STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
BUILD_TAGS = netgo

build:
	GOPRIVATE="github.com/opengovern" CC=/usr/bin/musl-gcc GOOS=linux GOARCH=amd64 go build -v -mod=mod -ldflags "-linkmode external -extldflags '-static' -s -w" -o ./build/steampipe-plugin-jira.plugin *.go

install:
	go build -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/turbot/jira@latest/steampipe-plugin-jira.plugin -tags "${BUILD_TAGS}" *.go
