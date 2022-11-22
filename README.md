# 🐮 mooxy 
Router for url paths, it uses a linked tree for path traversal to move fast trough large number of routes.

Instalation
``` 
go get github.com/mihai-craita/mooxy
```

Simple example
```
import "github.com/mihai-craita/mooxy"

func main() {
    router := mooxy.NewRouter()
    
    router.Handle(mooxy.NewRoute("/foo/bar"), handler)

    router.Handle(mooxy.NewRoute("/login").Methods("GET", "POST"), handler)
    
    imgRoute := mooxy.NewRoute("/img/{*}").Methods("GET")
    router.Handle(imgRoute, fileserver.GetFileServerHandler())
}
```
