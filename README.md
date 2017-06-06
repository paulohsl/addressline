Simple Free Text Address Parser
This is a simple Address Parser. It receives an simple String and returns two String fields:
1 - Street Name
2 - Street Number

Instructions:

Download Go. Download the latest version of Go for your platform here: https://golang.org/dl/.

Install Go. Follow the instructions for your platform to install the Go tools: https://golang.org/doc/install#install. It is recommended to use the default installation settings.

Test your Go installation.

Create and run the hello.go application described here: https://golang.org/doc/install#testing.

If you set up your Go environment correctly, you should be able to run “hello” from any directory and see the program execute successfully.

Installing Application

Follow the steps bellow to Install Package Application on your local machine

go get github.com/paulohsl/addressline/
cd $GOPATH/src/github.com/paulohsl/addressline/cmd/addressline
go install
Call compiled application and pass the log file as argument

addressline "Klosterstraße 62"

result: Klosterstraße,62

Running the tests

Navigate to root path application

cd $GOPATH/src/github.com/paulohsl/addressline
go test

PASS
ok  	github.com/paulohsl/webhookparser	0.007s
Authors

Paulo Henrique S. Lopes - Initial work -(https://github.com/paulohsl)
