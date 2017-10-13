## Week 5 Challenge
The code in the repository is incomplete and requires your help to put it back together! This will leverage the skills that you've learnt in the previous weeks and acts as a major checkpoint in preparation for the remaining weeks.

#### Description
We have a web server (code in server folder) that accepts messages sent to it at the `POST /messages` endpoint, and returns all messages so far at the `GET /echo` endpoint.

#### Ground work
Read through the source code to make sure you understand what the program aims to do. Explore any new packages you've not seen before yet to get a jist of what functionality the package offers.

#### Tasks
When you try to run `go build ./...` in the root folder, you'll see:
```
➜  echoserver git:(master) ✗ go build ./...
# upcode-academy-golang/echoserver/memory
memory/memory.go:5: imported and not used: "upcode-academy-golang/echoserver/interfaces"
memory/memory.go:22: missing return at end of function
# upcode-academy-golang/echoserver/filereader
filereader/filereader.go:7: imported and not used: "strings"
filereader/filereader.go:10: imported and not used: "upcode-academy-golang/echoserver/common"
filereader/filereader.go:13: undefined: interfaces in interfaces.DataProvider
filereader/filereader.go:22: unknown field 'FileName' in struct literal of type FileReader
filereader/filereader.go:29: f.FileName undefined (type *FileReader has no field or method FileName)
filereader/filereader.go:31: f.FileName undefined (type *FileReader has no field or method FileName)
```

Made a small change.

In its original state, the code will not compile. Fix up this program by looking in the source code for comments e.g `/* An interface method is missing here */`. The comment is a marker telling you that code needs to be added there for the program to work as expected.


#### How to run the server
`cd` to the server folder and run the `go run server.go` command. The server code requires no changes and you should see the following message on running the server:
```
➜  server git:(master) ✗ go run server.go
2017/10/07 10:55:04 Starting up the server!
```

#### Output
Once you've completed patching up the code and running both the server and main, you're all done if you see the exact output (except timestamps) when you visit http://localhost:8080/echo
```
2017-10-07 02:47:32.903994797 +0000 UTC: 
	********************
	Memory Buffer
	********************
	
2017-10-07 02:47:32.904353586 +0000 UTC: 5 affordable private islands in Asia
2017-10-07 02:47:32.904572627 +0000 UTC: The Asteria condo sells for $27.1 million
2017-10-07 02:47:32.904755573 +0000 UTC: iPhone 8 Plus batteries are swelling out of their cases
2017-10-07 02:47:32.904971362 +0000 UTC: 
	********************
	File Reader
	********************
	
2017-10-07 02:47:32.905123314 +0000 UTC: Always know what your code is supposed to do before you start writing it
2017-10-07 02:47:32.9052981 +0000 UTC: Test your code before you ship it.
2017-10-07 02:47:32.905525149 +0000 UTC: Practice makes perfect!
```