# ![HAProxy](../assets/images/haproxy-weblogo-210x49.png "HAProxy")

## HAProxy configuration parser

### example

```go
package main

import (
	"log"
)

func main() {
	p := Parser{}
	_, err := p.LoadData("/path/to/haproxy/file.cfg")
	log.Println(err)
	log.Println(p.String())

	data, err := p.GetAttr("global", "nbproc")
	if err == nil {
		//NOTE simple param may at some point get own structure
		nbproc := data.(*simple.SimpleNumber)
		log.Println(nbproc.Value)
	}	

	data, err = p.GetAttr("global", "stats socket")
	if err == nil {
		//NOTE simple param may at some point get own structure
		statsSocket := data.(*parsers.StatsSocket)
		log.Println(statsSocket.Path)
	}
}

```