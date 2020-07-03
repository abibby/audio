all: bin/sin bin/play bin/mul bin/add

bin/sin: sin/*.go *.go
	go build -o bin/sin ./sin

bin/play: play/*.go *.go
	go build -o bin/play ./play

bin/mul: mul/*.go *.go
	go build -o bin/mul ./mul

bin/add: add/*.go *.go
	go build -o bin/add ./add

clean:
	rm -rf bin/*
