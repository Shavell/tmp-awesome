````
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sql-connector main.go 
or
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v -ldflags "-s -w" -o main main.go
````