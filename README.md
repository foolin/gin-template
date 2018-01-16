# gin-template
golang template for gin framework!

# Feature
* Easy and simple to use for gin framework.
* Use golang html/template syntax.
* Support master layout file.
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

### Render master
```go
c.HTML(http.StatusOK, "index", gin.H{})
```

### Render file(not master layout)
```go
//use full name with extension
c.HTML(http.StatusOK, "page_file.tpl", gin.H{})
```


# Include
```html
<div>
{{include "layouts/footer"}}
</div>
```

# Usage
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

	router.HTMLRender = gintemplate.Default()
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"title": "Main website",
		})
	})

	//register router
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index", gin.H{
			"title": "Index title!",
			"escape": func(content string) string {
				return template.HTMLEscapeString(content)
			},
		})
	})

	router.GET("/page_file", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, page_file.tpl, gin.H{"title": "Page file title!!"})
	})

	router.Run(":9090")
}

```


# Advance
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
