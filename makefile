make:
	env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64  go build  -o build/project1.app/Contents/MacOS/project1
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -i -tags="tempdll" -o build/PycByte.exe
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -i -ldflags="-H windowsgui" -tags="tempdll" -o build/PycByte_CmdHide.exe
runmac:
	open build/project1.app
runwin:
	./build/PycByte.exe
