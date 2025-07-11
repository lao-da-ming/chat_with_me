# Proto 依赖目录配置，可以根据需求添加，按空格分割
# 注意：
# 1. bcl 为 bsi 内与 c# 项目共用的 proto 目录， 如果需要提供给 c# 项目使用，或者调用 C# 项目，需要添加到 proto_dirs 中
# 2. basic 为 bsi 的基础 proto 目录, 需要添加到 proto_dirs 中
# 3. validate 为 grpc validate 的依赖目录， 如果需要使用 grpc validate 需要添加到 proto_dirs 中，否则会报错
PROTO_DIRS := bcl basic panda validate axis/hrms hello  loris enums

# 项目名称 与main.go 上级目录名称保持一致
ProjectName = hrms

# 数据库名称
PgDbName = postgres
#mysql的 DbConnStr = account:password@tcp(ip:port)/$(DbName)
#测试数据库
PgDbConnStr =  user=postgres password=123456 dbname=$(PgDbName) host=10.60.33.25 port=5432 sslmode=disable TimeZone=Asia/Shanghai

# **********************************************************************************
# ********************************** 以下请不要修改 **********************************
# **********************************************************************************
GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

# export PROTO_DIRS for scripts
export PROTO_DIRS

#git config --global url."SYSTEM@git.shijizhongyun.com:".insteadOf "https://git.shijizhongyun.com/"

export GOPRIVATE=git.shijizhongyun.com/*

# check os
ifeq ($(OS),Windows_NT)
	PROTO_GEN := protoc.bat
else
	PROTO_GEN := ./protoc.sh
endif

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: gorm
#panda table entity generator
pgorm:# make pgorm or make pgorm table=users
ifdef table
	./gen-gorm --sqltype=postgres --connstr "$(PgDbConnStr)" --database $(PgDbName) --model entity --json --gorm --overwrite --out ./common/model --table $(table)
else
	./gen-gorm --sqltype=postgres --connstr "$(PgDbConnStr)" --database $(PgDbName) --model entity --json --gorm --overwrite --out ./common/model
endif
.PHONY: asset
asset:
	go test -v ./api/grpc/grpc_test.go -test.run TestParseAsset

.PHONY: build_arm
# build
build: asset
	mkdir -p bin/ && GOOS=linux GOARCH=arm64 go build -ldflags "-X main.Version=$(VERSION)" -o ./bin  ./...
	#dockerfile 不支持直接引用系统环境变量,更改CMD文件名称
	mv ./bin/$(ProjectName) ./bin/server



.PHONY: build
# build
build: asset
	mkdir -p bin/ && GOOS=linux go build -ldflags "-X main.Version=$(VERSION)" -o ./bin  ./...
	#dockerfile 不支持直接引用系统环境变量,更改CMD文件名称
	#mv ./bin/$(ProjectName) ./bin/server

# build
build-amd: asset
	mkdir -p bin/ && GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o ./bin  ./...

.PHONY: build-devops
# build
build-devops: build-amd
	mv ./bin/$(ProjectName) ./bin/server

.PHONY: upload
upload: build
	tool nexus upload --path=/$(ProjectName)/$(shell git branch --show-current) --asset=./bin/$(ProjectName)
	cd bin && rm $(ProjectName)

.PHONY: proto
proto:
	@if $(PROTO_GEN); then \
		echo "✅ protoc compilation succeeded"; \
	else \
		echo "❌ protoc compilation failed"; \
	fi

.PHONY: generate
# generate
generate:
	GOWORK=off go generate ./...
	go mod tidy

.PHONY: all
# generate all
all:
	make proto;
	make generate;

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
