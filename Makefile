tests:
	sudo GOGC=960 CGO_ENABLED=1 GOOS=linux GOARCH=386 go test ./... -v 

test:
	sudo GOGC=960 CGO_ENABLED=1 GOOS=linux GOARCH=386 go test

tests-win: 
	set GOGC=960
	set CGO_ENABLED=1
	set GOOS=linux
	set GOARCH=386
	go test ./... -v


test-cover:
	sudo GOGC=960 CGO_ENABLED=1 GOOS=linux GOARCH=386 go test -coverprofile=coverage.out ./...	
	go tool cover -html=coverage.out -o coverage.html

vet:
	sudo CGO_ENABLED=1 GOOS=linux GOARCH=386 go vet ./...
