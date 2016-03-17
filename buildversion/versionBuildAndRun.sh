GitHash=`git rev-parse HEAD`
BuildUser=`whoami`
BuildDay=`date "+%Y_%m_%d_%H:%M:%S"`
GitDetail=`git log -1`

Msg="Binay Build By $BuildUser   at $BuildDay  with verion $GitHash"
#Msg="$Msg /n Last commit is"
#Msg="$Msg /n $GitDetail"


CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version '$Msg' "
echo ""
echo "--prepare to run--"
echo ""
./buildversion