# Go Profiling and Optimization

This code is example code used in the talk [Go Profiling and Optimization](https://docs.google.com/presentation/d/1n6bse0JifemG7yve0Bb0ZAC-IWhTQjCNAclblnn2ANY/edit#slide=id.g3a3e2af65_029).

It shows how pprof and [go-torch](https://github.com/uber/go-torch) can be
used to identify performance bottlenecks, and optimize them.




go run main.go -printStats



go run main.go
go get github.com/adjust/go-wrk
go-wrk -d 5 http://localhost:9090/simple


http://brew.sh/index_zh-cn.html
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
brew install graphviz
go tool pprof --seconds 5 http://localhost:9090/debug/pprof/heap?debug=1
> top 10
> web

go tool pprof --seconds 5 http://localhost:9090/debug/pprof/profile
> web


git clone  https://github.com/brendangregg/FlameGraph
go-torch --seconds 5 http://localhost:9090/debug/pprof/profile