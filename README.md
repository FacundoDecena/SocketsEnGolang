# Socket in Golang

to be honest I'm not sure if this are sockets or not...

## for running server

```bash
 ./build/server
 ```

## for running client

```bash
./build/env
```

What? Env is not there?

just make a file named env.go in /client and insert this code

```Golang
package main

const IP string = "server ip" // 192.168.0.x

func getIp() string{
 return IP
}
```

then go build -o build/ client/env.go go build -0 build/ client/env.go

and you are good to go
