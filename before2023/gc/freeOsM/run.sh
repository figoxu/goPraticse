go fmt && go install && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build
export GOGC=30
GODEBUG=gctrace=1 ./freeOsM