whling/go-config is a Go package that provides a generic configuration interface which enables a Golang application to read configuration data from a variety of sources.

Getting Started
===============

## Installing

To start using whling/go-config, install Go and run `go get`:

```sh
$ go get -u github.com/whling/go-config
```

## Usage example
Read a TOML document:
```go
import "github.com/whling/go-config"

func main() {

    // LoadTomlFile() creates a Node from a file.
    node, _ := config.LoadTomlFile("server.toml")
    
    // LoadToml() creates a Node from a string.
    node, _ = config.LoadToml(`
    [server]
    name = "nginx"
    log = "var/log/access.log"
    port = 8080
    timeout = "100ms"
    ip = ["127.0.0.1", "192.168.1.1", "10.123.234.2"]`)
    
    // Init
    cfg := &config.Value{node}
    
    // Getter
    name := cfg.Str("server.name", "none")
    port := cfg.Int64("server.port", 80)
    timeout := cfg.Duration("server.timeout", 1*time.Second)
    ip := cfg.StrArray("server.ip")
    log := cfg.Access("server").Access("log").Value().(string)

}
```
