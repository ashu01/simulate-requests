About
=====
Repo displays a simple HTTP request-response, using JSON objects. It can be used to simulate a client request (from Python, using the requests module) and a corresponding server response (from Golang, using the gin module).

Instructions
============
1. Set up Python 3 and pip3
2. Install requests module using the following command `sudo pip3 install requests`
3. Install gin using the following command `go get github.com/gin-gonic/gin`
4. Clone the repo
5. On one terminal window, run `go run server.go`
6. On another terminal window, run `python3 send.py`
7. The request should be visible on server.go and the corresponding response should be visible on the terminal running the python program
