GO_CMD=go
GO_BUILD=${GO_CMD} build

SERVER_BIN=bin_terraasri_code_challenge
SERVER_DIR=.
SERVER_MAIN=main.go

build:
	@${GO_BUILD} -o ${SERVER_BIN} -ldflags "-w -s" ${SERVER_DIR}/${SERVER_MAIN} || exit1;
	@chmod 755 ${SERVER_BIN}

run:
	./${SERVER_BIN}