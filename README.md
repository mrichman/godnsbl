# godnsbl [![Travis-CI](https://travis-ci.org/mrichman/godnsbl.svg)](https://travis-ci.org/mrichman/godnsbl) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/mrichman/godnsbl)](https://goreportcard.com/report/github.com/mrichman/godnsbl)

Package godnsbl lets you perform RBL (Real-time Blackhole List - https://en.wikipedia.org/wiki/DNSBL)
lookups using Golang

The command-line tool in `cmd` demonstrates the use of [goroutines](https://tour.golang.org/concurrency/1) to perform concurrent lookups.

To test:

```
go get github.com/mrichman/godnsbl
cd $GOPATH/src/github.com/mrichman/godnsbl/cmd/godnsbl
go run main.go 127.0.0.2
```

The output will be a JSON-formatted list of results with the following fields:

```
[
...
{  
   "address":"127.0.0.2",
   "listed":true,
   "text":"Client host blocked using Barracuda Reputation, see http://www.barracudanetworks.com/reputation/?r=1\u0026ip=127.0.0.2",
   "error":false,
   "error_type":null
}
...
]
```
