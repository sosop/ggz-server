APP = GGZ

build:
	go build -o ./bin/${APP} -ldflags '-s -w'

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/${APP} -ldflags '-s -w'

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/win-${APP}.exe -ldflags '-s -w'

mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./bin/mac-${APP} -ldflags '-s -w'

freebsd:
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o ./bin/fbd-${APP} -ldflags '-s -w'

run:
	@go run *.go

clean:
	@rm ./bin/${APP}
