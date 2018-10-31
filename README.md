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
}

```