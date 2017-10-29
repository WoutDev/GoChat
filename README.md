# GoChat
Simple CLI chat application written in Go

## About  
This application is my first application written in Go. The goal was to use my newly acquired Go training and create a simple CLI application with it. The source code in my opinion is quite a mess tho.

## Building the server and client  
Just make sure you have a correct installation of Go on your system to be able to build the native binaries. 
```
git clone https://github.com/WoutDev/GoChat
cd GoChat
go get ./...
```

### Server
```
cd src/github.com/woutdev/gochat/server
go build server.go
```

### Client
(navigate back to the project root if you just builded the server)
```
cd src/github.com/woutdev/gochat/client
go build client.go
```

## Use
### Server
![server use](https://i.gyazo.com/d6b720579f73010b403aa061fcfab484.png)

### Client
![client use](https://i.gyazo.com/c5781093480ee34117f34bfc64382b1d.png)
