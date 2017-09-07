About
=====
Repo displays a simple HTTP request-response, using JSON objects. It can be used to simulate a client request (from Python, using the requests module) and a corresponding server response (from Golang, using the gin module).

Instructions
=============
1. Install and set up `go` && `postgres psql` 
2. Set up Python3.4 and pip3
3. Install requests module using the following command `sudo pip3 install requests`
4. Install gin using the following command `go get github.com/gin-gonic/gin`
5. Install postgres psql package using command `go get github.com/jackc/pgx`
6. Install this module for http requests `go get github.com/parnurzeal/gorequest`
7. Clone the repo
8. On one terminal window (say t1), run `go run server.go`
9. On another terminal window (say t2), run `cd API`  
10. run `create_table.py`   deletes the existing entries of `users` and `usersdescription` tables if table exists , else create tables
11. run `createuserid.py`   check and validate an `email` with it's userid
12. run `dumpintousers.py`  dump data into `users` table
13. run `dumpintousersdescription.py` dump data into `usersdescription` table
14. run `findDeviceidPlatform.py`     it finds the `deviceid` and `platform`
11. On another terminal window, run `python3.4 findDeviceidPlatform.py`
12. The request should be visible on server.go and the corresponding response should be visible on the terminal running the python program

Other python Modules used
=========================
  random, psycopg2, json
  install python modules `sudo pip3 install <package>`


Contributors
============
1. Ashutosh Kumar Gupta


[Doc](https://ashu01.github.io/simulate-requests/)
