# Basic
Support template for [go.rice](https://github.com/GeertJohan/go.rice)


# Install
```go
go get github.com/foolin/gin-template/supports/gorice
```

# Useage

```go
 gin.Renderer = gorice.New(rice.MustFindBox("views"))
```

# Example
```go

func main() {

	// Echo instance
	e := gin.New()

	// servers other static files
	staticBox := rice.MustFindBox("static")
	staticFileServer := http.StripPrefix("/static/", http.FileServer(staticBox.HTTPBox()))
	e.GET("/static/*", gin.WrapHandler(staticFileServer))

	//Set Renderer
	e.Renderer = gorice.New(rice.MustFindBox("views"))

	// Start server
	e.Logger.Fatal(e.Start(":9090"))
}

```

[gorice example](https://github.com/foolin/gin-template/tree/master/examples/gorice)

# Links

- [gin template](https://github.com/foolin/gin-template)
- [gin framework](https://github.com/gin-gonic/gin)
- [go.rice](https://github.com/GeertJohan/go.rice)
