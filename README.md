# ![HAProxy](../assets/images/haproxy-weblogo-210x49.png "HAProxy")

## HAProxy configuration parser

### example

```go
package main

//...

func main() {
    p := Parser{}
    _, err := p.LoadData("/path/to/haproxy/file.cfg")
    log.Println(err)
    log.Println(p.String())

    data, err := p.GetGlobalAttr("nbproc")
    if err == nil {
        //NOTE simple param may at some point get own structure
        nbproc := data.(*simple.SimpleNumber)
        log.Println(nbproc.Value)
    }

    data, err = p.GetGlobalAttr("stats socket")
    if err == nil {
        statsSocket := data.(*stats.Socket)
        log.Println(statsSocket.Path)
        statsSocket.Path = "/var/run/haproxy-runtime-api.0.sock"
    }

    data, err = p.GetDefaultsAttr("timeout http-request")
    if err == nil {
        //NOTE simple param may at some point get own structure
        statsSocket := data.(*simple.SimpleTimeout)
        log.Println(statsSocket.Value)
    }
    p.Save(configFilename)
}

```