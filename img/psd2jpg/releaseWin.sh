rsrc -manifest main.exe.manifest -ico rc.ico -o psd2jpg.syso
GOOS=windows GOARCH=amd64 go build -o psd2jpg.exe github.com/figoxu/goPraticse/img/psd2jpg