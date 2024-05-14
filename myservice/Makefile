.PHONY: build clean deploy

set-environment:
	set GOOS=linux                                             
	set GOARCH=arm64                                           
	set CGO_ENABLED=0        
zipfile:
	cd bin/postEmployee && 7z a ../postEmployee.zip bootstrap
	
build:
	go build -o bin/postEmployee/bootstrap functions/postEmployee/main.go   
	