````
GOOS=linux GOARCH=amd64 go build main.go
or
GOOS=linux GOARCH=arm64 go build -v -ldflags "-s -w" -o main main.go
````