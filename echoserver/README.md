## Week 5 Challenge
The code in the repository is incomplete and requires your help to put it back together! This will leverage the skills that you've learnt in the previous weeks and acts as a major checkpoint in preparation for the remaining weeks.

#### Description
We have a web server (code in server folder) that accepts messages sent to it at the `POST /messages` endpoint, and returns all messages so far at the `GET /echo` endpoint.

#### Ground work
Read through the source code to make sure you understand what the program aims to do. Explore any new packages you've not seen before yet to get a jist of what functionality the package offers.

#### Tasks
In its original state, the code will not compile. Fix up this program by looking in the source code for comments e.g `/* An interface method is missing here */`. The comment is a marker telling you that code needs to be added there for the program to work as expected.

#### How to run the server
`cd` to the server folder and run the `go run server.go` command. The server code requires no changes and you should see the following message on running the server:
```
➜  server git:(master) ✗ go run server.go
2017/10/07 10:55:04 Starting up the server!
```
