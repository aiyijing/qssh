build:
	go build -o bin/qssh main.go

install: build
	install ./bin/qssh /usr/local/bin/

clean:
	rm -rf bin/qssh