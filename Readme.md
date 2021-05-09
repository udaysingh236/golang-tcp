This is a TCP server in golang.
The server is designed on Job/Worker Pattern by using channel and go routines. Server will register workers provided in the "-worker" flag as it starts. Dispatcher will dispatch the work request to the availble worker.

For now the server is listening 10 bytes at a time, you can change it in the readNextByte function of collectandvalidate.go module and build it again.

To Build:
go build -o 'tcp-worker' *.go

Usage:
./tcp-worker --help
Usage of D:\GOPO\tcp-worker:
  -worker int
        The number of workers you want to run..! (default 4)

./tcp-worker -worker 100

App is structured in four goroutines to perform four basic tasks of TCP server

1.       Read from TCP port, break input stream into frames
2.       Dispatch one request message at a time to a worker
3.       Process request message (Business logic)
4.       Write to TCP port
![app structure image](https://github.com/aniketchopade/golang-tcp/blob/master/app-structure.png)
