# Notes

## Simple TCP server and client
* A TCP server waits before writing back to a client whereas a client
  writes first and then waits for a response.
* Installation:

```
//Start TCP server on port 7777 on localhost
go install ./cmd/tcpserver;tcpserver

//Start TCP Client
go install ./cmd/tcpclient;client 127.0.0.1:7777
```

* The client and server can also be ran on different IP addresses so
  long as the TCP client can make a call out the TCP server IP:port

## Simple UDP server and client
```
//Start UDP server on port 7777 on localhost
go install ./cmd/udpserver;udpserver 127.0.0.1:7777

//Start UDP Client
go install ./cmd/client;client 127.0.0.1:7777
```
