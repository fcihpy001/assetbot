
.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/assetBot ../main.go

.PHONY: mac
mac:
	go build -o ./build/manager main.go

# 编译到 windows
.PHONY: windows
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./build/assetBot.exe main.go

.PHONY: clean
clean:
	rm -rf ./build

# 编译到 全部平台
.PHONY: build-all
all:
	make clean
	#make mac
	make linux




