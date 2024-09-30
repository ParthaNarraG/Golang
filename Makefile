
#If there is no target mentioned in the command, this will act as default target($make)
.DEFAULT_GOAL := run

.PHONY:
	fmt vet build

#Formats the go code
fmt:
	go fmt ./...

#Founds any errors
vet: fmt 
	go vet ./...

#creates a binary file which can execute go code
build: vet 
	go build

#runs your go program
run: 
	go run first-program.go

#Removes the binary files created from running go build
clean:
	go clean