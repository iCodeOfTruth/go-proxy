# go-list
Proxy Manager created in Go

# How to install
```
go get -u github.com/iCodeOfTruth/go-proxy
```

# Usage

```golang
package main

import goproxy "github.com/iCodeOfTruth/go-proxy"

var Proxies = []interface{}{"ip:port", "ip:port:username:password",
	"username:password@ip:port"}

func main() {
	proxyM := goproxy.NewProxyManager(Proxies)

	// Returns proxy and error
	proxyM.GetNextProxy() // Get Next Proxy
	proxyM.GetRandomProxy() // Get Random Proxy
}
```