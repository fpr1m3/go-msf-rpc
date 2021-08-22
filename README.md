# msf-rpc-client
Golang based RPC client to communicate with Metasploit. This is based off Black Hat Go's [example](https://github.com/blackhat-go/bhg/tree/master/ch-3/metasploit-minimal). Extended to include other methods from the RPC API by @fpr1m3.

This modification allows for commands to be run on existing sessions.

## Install
`go get -u github.com/fpr1m3/msf-rpc-client/...`

## Use
`export MSFHOST=127.0.0.1:55552`

`export MSFPASS=password`

`go run client/main.go`

`go run client/main.go commands.txt`

The additional parameter is a list of commands to be run delimited by newline characters:
Ex: 
```
whoami
id
hostname
```
### Reference
https://wya.pl/2020/04/27/metasploits-rpc-api/
