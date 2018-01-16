# gin-template
golang template for gin framework!

# Feature
* Easy and simple to use for gin framework.
* Use golang html/template syntax.
* Support configure master layout file.
* Support configure template file extension.
* Support configure templates directory.
* Support configure cache template.
* Support include file.

# Configure

```go
    gintemplate.TemplateConfig{
		Root:      "views",
		Extension: ".tpl",
		Master:    "layouts/master",
		Partials:  []string{"partials/head"},
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
		},
		DisableCache: false,
	}
```


# Render

### Render with master
```go
//use name without extension
c.HTML(http.StatusOK, "index", gin.H{})
```

### Render only file(not use master layout)
```go
//use full name with extension
c.HTML(http.StatusOK, "page_file.tpl", gin.H{})
```


# Include syntax
```html
<div>
{{include "layouts/footer"}}
</div>
```

# Usage

## Install
```bash
go get github.com/foolin/gin-template
```

## Basic example
```go

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/foolin/gin-template"
	"net/http"
)

func main() {
	router := gin.Default()

	//new template engine
	router.HTMLRender = gintemplate.Default()

	router.GET("/", func(ctx *gin.Context) {
		//render with master
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})


	router.GET("/page_file", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "page_file.html", gin.H{"title": "Page file title!!"})
	})

	router.Run(":9090")
}


```
[Basic example](https://github.com/foolin/gin-template/tree/master/examples/basic)

## Advance example
```go

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/foolin/gin-template"
	"net/http"
	"html/template"
)

func main() {
	router := gin.Default()

	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:      "views",
		Extension: ".tpl",
		Master:    "layouts/master",
		Partials:  []string{"partials/head"},
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
		},
		DisableCache: false,
	})

	router.GET("/", func(ctx *gin.Context) {
		//render with master
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})

	router.GET("/page_file", func(ctx *gin.Context) {
		//render only file, must full name with extension
		ctx.HTML(http.StatusOK, "page_file.tpl", gin.H{"title": "Page file title!!"})
	})

	router.Run(":9090")
}

```
[Advance example](https://github.com/foolin/gin-template/tree/master/examples/advance)
