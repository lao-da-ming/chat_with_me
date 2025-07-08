
# 数据库名称
DbName = bsi_axis_hrms
PgDbName = bsi_panda
# 数据库连接字符串
DbConnStr = sjzy_dev_user:818f83e5@tcp(mysql-dev.default.svc.cluster.local:3306)/$(DbName)
#磐石数据库
PgDbConnStr =  user=huangdong password=66881234zfkaka7MM_MzX dbname=$(PgDbName) host=common-pgdb.dev.sjzy.local port=5432 sslmode=disable TimeZone=Asia/Shanghai
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
.PHONY: gorm
gorm: # make gorm or make gorm table=users
ifdef table
	gen-gorm --sqltype=mysql --connstr "$(DbConnStr)" --database $(DbName) --model entity --json --gorm --overwrite --out ./common/model --table $(table)
else
	gen-gorm --sqltype=mysql --connstr "$(DbConnStr)" --database $(DbName) --model entity --json --gorm --overwrite --out ./common/model
endif
.PHONY: gorm
#panda table entity generator
pgorm:# make pgorm or make pgorm table=users
ifdef table
	gen-gorm --sqltype=postgres --connstr "$(PgDbConnStr)" --database $(PgDbName) --model entity --json --gorm --overwrite --out ./common/model --table $(table)
else
	gen-gorm --sqltype=postgres --connstr "$(PgDbConnStr)" --database $(PgDbName) --model entity --json --gorm --overwrite --out ./common/model
endif