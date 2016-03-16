hcode=`git rev-parse HEAD`
u=`whoami`
bd=`date "+%Y_%m_%d_%H:%M:%S"`

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version $u$bd$hcode"
echo ""
echo "--prepare to run--"
echo ""
./buildversion