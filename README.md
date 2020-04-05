# Socket in Golang

to be honest I'm not sure if this are sockets or not...

## for running server

./build/server

## for running client

./build/env

What? Env is not there?

just make a file named env.go and insert this code
```package main

const IP string = "server ip" // 192.168.0.x

func getIp() string{
	return IP
}```

the go build -0 build/ client/env.go

and you are good to go
