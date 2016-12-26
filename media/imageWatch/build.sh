#GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version '[version] Binary Build By `whoami` at `date "+%Y_%m_%d_%H:%M:%S"` with version `git rev-parse HEAD`' "

GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version '[version] Binary Build By `whoami` at `date "+%Y_%m_%d_%H:%M:%S"` with version `git rev-parse HEAD`' "