# go-domoto

go-domoto is a Golang library for accessing the Domoticz HTTP API.

Usage is simple:

```golang
c := domoto.New("http://localhost:8080","username","password"
res, err := c.AllDevices("")
```
or
```golang
res, err := c.DeviceToggle(1)
```

A more complete example is an implementation of a CLI here:
https://github.com/pawal/domo-cli
